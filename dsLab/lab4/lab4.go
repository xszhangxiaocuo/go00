package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*
bmp图像处理，顺时针旋转九十度，逆时针旋转九十度，水平翻转，垂直翻转
2023年4月14日10:19:18
*/

type bmpFileHeader struct {
	FileType     [2]byte //2 字节，用于存储文件类型。对于 BMP 文件，这个字段的值应该是 "BM"
	FileSize     uint32  //4 字节，用于存储整个 BMP 文件的大小（以字节为单位）
	Reserved1    uint16  //2 字节，保留字段，通常设置为 0。应用程序可以使用这个字段来存储自定义信息。
	Reserved2    uint16  //2 字节，保留字段，通常设置为 0。应用程序可以使用这个字段来存储自定义信息。
	BitmapOffset uint32  //4 字节，用于存储从文件开始到实际位图数据开始之间的偏移量（以字节为单位）。
}

type bmpInfoHeader struct {
	Size            uint32 //4 字节，用于存储 bmpInfoHeader 结构体的大小（以字节为单位）。对于 Windows BMP 文件，这个值通常是 40。
	Width           int32  //4 字节，用于存储图像的宽度（以像素为单位）。
	Height          int32  //4 字节，用于存储图像的高度（以像素为单位）。
	Planes          uint16 //2 字节，用于存储颜色平面数。对于 BMP 文件，这个值应该是 1。
	BitsPerPixel    uint16 //2 字节，用于存储每个像素的颜色深度（以位为单位）。常见的值有 1（黑白图像，单色），4（16 色），8（256 色），16（65536 色），24（16777216 色，真彩色）和 32（4294967296 色，带有 Alpha 通道的真彩色）。
	Compression     uint32 //4 字节，用于存储图像数据的压缩类型。常见的值有 0（BI_RGB，无压缩），1（BI_RLE8，RLE 压缩，仅用于 8 位/像素位图），2（BI_RLE4，RLE 压缩，仅用于 4 位/像素位图）和 3（BI_BITFIELDS，无压缩，带有颜色掩码，仅用于 16 和 32 位/像素位图）。
	SizeOfBitmap    uint32 //4 字节，用于存储位图数据的大小（以字节为单位）。如果图像没有压缩，这个值可以设置为 0。
	HorzResolution  int32  //4 字节，用于存储图像的水平分辨率（以像素/米为单位）。通常这个值可以设置为 0。
	VertResolution  int32  //4 字节，用于存储图像的垂直分辨率（以像素/米为单位）。通常这个值可以设置为 0。
	ColorsUsed      uint32 //4 字节，用于存储位图中实际使用的颜色数。对于 8 位/像素及以下的位图，这个值表示调色板中的颜色数。如果使用了所有可能的颜色（即调色板大小为 2 的 BitsPerPixel 次方），这个值可以设置为 0。
	ColorsImportant uint32 //4 字节，用于存储对图像显示有重要影响的颜色数。通常，这个值可以设置为 0，表示所有颜色都是重要的。
}

func main() {
	workDir, _ := os.Getwd()
	fileHeader, infoHeader, palette, bitmapData, err := read(workDir + "/lab4/flowers.bmp")
	if err != nil {
		fmt.Println("Error rotating BMP file:", err)
	}
	// 顺时针
	rotatedBitmapData := leftBitmapData(bitmapData, int(infoHeader.Width), int(infoHeader.Height), int(infoHeader.BitsPerPixel))
	err = save(workDir+"/lab4/left.bmp", *fileHeader, *infoHeader, palette, rotatedBitmapData, true)
	if err != nil {
		fmt.Println("Error write BMP file:", err)
	}
	fmt.Println("success!")
	//逆时针
	rotatedBitmapData = rightBitmapData(bitmapData, int(infoHeader.Width), int(infoHeader.Height), int(infoHeader.BitsPerPixel))
	err = save(workDir+"/lab4/right.bmp", *fileHeader, *infoHeader, palette, rotatedBitmapData, true)
	if err != nil {
		fmt.Println("Error write BMP file:", err)
	}
	fmt.Println("success!")
	//水平翻转
	rotatedBitmapData = FlipHBitmapData(bitmapData, int(infoHeader.Width), int(infoHeader.Height), int(infoHeader.BitsPerPixel))
	err = save(workDir+"/lab4/fliph.bmp", *fileHeader, *infoHeader, palette, rotatedBitmapData, false)
	if err != nil {
		fmt.Println("Error write BMP file:", err)
	}
	fmt.Println("success!")
	//垂直翻转
	rotatedBitmapData = FlipVBitmapData(bitmapData, int(infoHeader.Width), int(infoHeader.Height), int(infoHeader.BitsPerPixel))
	err = save(workDir+"/lab4/flipv.bmp", *fileHeader, *infoHeader, palette, rotatedBitmapData, false)
	if err != nil {
		fmt.Println("Error write BMP file:", err)
	}
	fmt.Println("success!")
}

func read(source string) (*bmpFileHeader, *bmpInfoHeader, []byte, [][]byte, error) {
	// 打开源文件
	srcFile, err := os.Open(source)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error opening source file: %w", err)
	}
	defer srcFile.Close()

	// 读取 BMP 头
	fileHeader := bmpFileHeader{}
	err = binary.Read(srcFile, binary.LittleEndian, &fileHeader)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error reading file header: %w", err)
	}

	infoHeader := bmpInfoHeader{}
	err = binary.Read(srcFile, binary.LittleEndian, &infoHeader)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error reading info header: %w", err)
	}

	if fileHeader.FileType != [2]byte{'B', 'M'} {
		return nil, nil, nil, nil, fmt.Errorf("invalid BMP file: not a valid BMP file")
	}

	// 校验是否为 8 位 BMP 文件
	if infoHeader.BitsPerPixel != 8 {
		return nil, nil, nil, nil, fmt.Errorf("unsupported BMP format: only 8-bit BMP files are supported")
	}

	// 读取调色板
	paletteSize := 1 << infoHeader.BitsPerPixel
	palette := make([]byte, paletteSize*4)
	_, err = srcFile.Read(palette)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("error reading palette: %w", err)
	}
	//将文件指针移到实际位图数据的开始地址
	srcFile.Seek(int64(fileHeader.BitmapOffset), 0)

	width := int(infoHeader.Width)
	height := int(infoHeader.Height)

	// 读取位图数据
	bitmapData := make([][]byte, height)
	paddedWidth := (((width*int(infoHeader.BitsPerPixel))/8 + 3) / 4) * 4
	for i := range bitmapData {
		bitmapData[i] = make([]byte, paddedWidth)
		_, err := srcFile.Read(bitmapData[i])
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("error reading bitmap data: %w", err)
		}
	}
	return &fileHeader, &infoHeader, palette, bitmapData, nil
}

func save(destination string, fileHeader bmpFileHeader, infoHeader bmpInfoHeader, palette []byte, rotatedBitmapData [][]byte, flag bool) error {
	// 创建并写入目标文件
	dstFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer dstFile.Close()

	// 如果flag为true，交换宽度和高度
	if flag {
		infoHeader.Width, infoHeader.Height = infoHeader.Height, infoHeader.Width
	}

	// 将文件头和信息头写入目标文件
	err = binary.Write(dstFile, binary.LittleEndian, &fileHeader)
	if err != nil {
		return fmt.Errorf("error writing file header: %w", err)
	}
	err = binary.Write(dstFile, binary.LittleEndian, &infoHeader)
	if err != nil {
		return fmt.Errorf("error writing info header: %w", err)
	}

	// 写入调色板
	_, err = dstFile.Write(palette)
	if err != nil {
		return fmt.Errorf("error writing palette: %w", err)
	}

	// 将旋转后的位图数据写入目标文件
	for _, row := range rotatedBitmapData {
		_, err = dstFile.Write(row)
		if err != nil {
			return fmt.Errorf("error writing rotated bitmap data: %w", err)
		}
	}
	return nil
}

// 顺时针旋转九十度
func leftBitmapData(bitmapData [][]byte, width, height int, bitsPerPixel int) [][]byte {
	rotatedBitmapData := make([][]byte, width)
	paddedRotatedWidth := (((height*bitsPerPixel)/8 + 3) / 4) * 4
	for i := range rotatedBitmapData {
		rotatedBitmapData[i] = make([]byte, paddedRotatedWidth)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			newX, newY := y, x
			rotatedBitmapData[x][y] = bitmapData[newX][newY]
		}
	}

	return rotatedBitmapData
}

// 逆时针旋转九十度
func rightBitmapData(bitmapData [][]byte, width, height int, bitsPerPixel int) [][]byte {
	rotatedBitmapData := make([][]byte, width)
	paddedRotatedWidth := (((height*bitsPerPixel)/8 + 3) / 4) * 4
	for i := range rotatedBitmapData {
		rotatedBitmapData[i] = make([]byte, paddedRotatedWidth)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			newX, newY := height-y-1, x
			rotatedBitmapData[x][y] = bitmapData[newX][newY]
		}
	}

	return rotatedBitmapData
}

// 水平翻转
func FlipHBitmapData(bitmapData [][]byte, width, height int, bitsPerPixel int) [][]byte {
	rotatedBitmapData := make([][]byte, height)
	paddedRotatedWidth := (((width*bitsPerPixel)/8 + 3) / 4) * 4
	for i := range rotatedBitmapData {
		rotatedBitmapData[i] = make([]byte, paddedRotatedWidth)
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			newX, newY := x, width-y-1
			rotatedBitmapData[x][y] = bitmapData[newX][newY]
		}
	}

	return rotatedBitmapData
}

// 垂直翻转
func FlipVBitmapData(bitmapData [][]byte, width, height int, bitsPerPixel int) [][]byte {
	rotatedBitmapData := make([][]byte, height)
	paddedRotatedWidth := (((width*bitsPerPixel)/8 + 3) / 4) * 4
	for i := range rotatedBitmapData {
		rotatedBitmapData[i] = make([]byte, paddedRotatedWidth)
	}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			newX, newY := height-x-1, y
			rotatedBitmapData[x][y] = bitmapData[newX][newY]
		}
	}

	return rotatedBitmapData
}

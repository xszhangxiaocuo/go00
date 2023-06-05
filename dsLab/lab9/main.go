package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

/*
排序算法性能比较试验
2023年6月5日14:22:19
*/

type data struct {
	id    int
	value int
}

var compare = 0
var exchange = 0
var t float64

func main() {
	index := "1"
	num := "50000"
	file := "lab9/data/" + num + "_" + index + ".xlsx"
	output := []string{
		"lab9/sort/siSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/shellSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/bubbleSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/quickSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/selectSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/heapSort_" + num + "_" + index + "_output.xlsx",
		"lab9/sort/mergeSort_" + num + "_" + index + "_output.xlsx",
	}
	d := readData(file)
	siSort(d)
	writeFile(d, output[0])
	fmt.Printf("直接插入排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	shellSort(d)
	writeFile(d, output[1])
	fmt.Printf("希尔排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	bubbleSort(d)
	writeFile(d, output[2])
	fmt.Printf("冒泡排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	start := time.Now()
	compare = 0
	exchange = 0
	quickSort(d, 0, len(d)-1)
	t = time.Since(start).Seconds()
	writeFile(d, output[3])
	fmt.Printf("快速排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	selectSort(d)
	writeFile(d, output[4])
	fmt.Printf("简单选择排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	start = time.Now()
	compare = 0
	exchange = 0
	heapSort(d)
	t = time.Since(start).Seconds()
	writeFile(d, output[5])
	fmt.Printf("堆排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)

	d = readData(file)
	start = time.Now()
	compare = 0
	exchange = 0
	mergeSort(d)
	t = time.Since(start).Seconds()
	writeFile(d, output[6])
	fmt.Printf("归并排序[比较次数：%d,交换次数：%d,排序时间：%fs]\n", compare, exchange, t)
}

// 读取数据
func readData(file string) []data {
	excelFileName := file
	d := make([]data, 0)
	// 打开XLSX文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开文件：", err)
		return nil
	}

	// 遍历所有工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历每一行
		for _, row := range sheet.Rows {
			tmp := data{}
			// 遍历每个单元格
			for col, cell := range row.Cells {
				value, _ := cell.Int()
				switch col {
				case 0:
					tmp.id = value
				case 1:
					tmp.value = value
				}
			}
			d = append(d, tmp)
		}
	}
	return d
}

// 输出排序结果
func writeFile(d []data, outputName string) {
	// 创建一个新的xlsx文件
	file := xlsx.NewFile()

	// 创建一个名为"Sheet1"的工作表
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println("创建工作表时出错:", err)
		return
	}
	// 遍历结构体切片，将数据写入工作表
	for _, tmp := range d {
		row := sheet.AddRow()
		idCell := row.AddCell()
		idCell.SetInt(tmp.id)

		valueCell := row.AddCell()
		valueCell.SetInt(tmp.value)
	}

	// 保存文件
	err = file.Save(outputName)
	if err != nil {
		fmt.Println("保存XLSX文件时出错:", err)
		return
	}
}

// 直接插入排序
func siSort(d []data) {
	start := time.Now()
	compare = 0
	exchange = 0
	l := len(d)
	for i := 1; i < l; i++ {
		key := d[i]
		j := i - 1

		// 将比key大的元素向右移动
		for j >= 0 && d[j].value > key.value {
			if d[j].value > key.value {
				exchange++
			}
			compare++
			d[j+1] = d[j]
			j--
		}

		d[j+1] = key
	}
	t = time.Since(start).Seconds()
}

// 希尔排序
func shellSort(d []data) {
	start := time.Now()
	compare = 0
	exchange = 0
	l := len(d)
	gap := l / 2

	// 逐步减小间隔直到为1
	for gap > 0 {
		// 对每个间隔进行插入排序
		for i := gap; i < l; i++ {
			temp := d[i]
			j := i

			// 将比temp大的元素向右移动
			for j >= gap && d[j-gap].value > temp.value {
				compare++
				if d[j-gap].value > temp.value {
					exchange++
				}
				d[j] = d[j-gap]
				j -= gap
			}

			d[j] = temp
		}

		gap /= 2
	}
	t = time.Since(start).Seconds()
}

// 冒泡排序
func bubbleSort(d []data) {
	start := time.Now()
	compare = 0
	exchange = 0
	l := len(d)
	tmpd := d
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			compare++
			if tmpd[j].value > tmpd[j+1].value {
				tmpd[j], tmpd[j+1] = tmpd[j+1], tmpd[j]
				exchange++
			}
		}
	}
	t = time.Since(start).Seconds()
}

// 快速排序
func quickSort(d []data, left int, right int) {
	if left < right {
		pivotIndex := partition(d, left, right) // 获取划分点的索引
		quickSort(d, left, pivotIndex-1)        // 对划分点左侧子数组进行快速排序
		quickSort(d, pivotIndex+1, right)       // 对划分点右侧子数组进行快速排序
	}
}

// 划分函数，将数组划分成两个子数组，并返回划分点的索引
func partition(d []data, left, right int) int {
	pivot := d[right] // 选择最后一个元素作为划分点
	i := left - 1     // 划分点的索引，初始值为low-1

	for j := left; j < right; j++ {
		compare++
		if d[j].value < pivot.value {
			i++
			d[i], d[j] = d[j], d[i] // 将小于划分点的元素交换到左侧子数组中
		}
	}

	d[i+1], d[right] = d[right], d[i+1] // 将划分点元素放置到正确的位置上
	exchange++
	return i + 1 // 返回划分点的索引
}

// 简单选择排序
func selectSort(d []data) {
	start := time.Now()
	compare = 0
	exchange = 0
	l := len(d)
	for i := 0; i < l; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			compare++
			if d[j].value < d[k].value {
				k = j
			}
		}
		if k != i {
			d[i], d[k] = d[k], d[i]
			exchange++
		}
	}

	t = time.Since(start).Seconds()
}

// 堆排序函数
func heapSort(d []data) {
	n := len(d)

	// 构建最大堆
	buildMaxHeap(d)

	// 逐步将最大元素移到堆尾，然后调整堆
	for i := n - 1; i > 0; i-- {
		d[0], d[i] = d[i], d[0] // 将堆顶元素与堆尾元素交换
		exchange++

		// 调整堆
		heapify(d, 0, i)
	}
}

// 构建最大堆
func buildMaxHeap(d []data) {
	n := len(d)

	// 从最后一个非叶子节点开始，向上调整每个子树使其成为最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(d, i, n)
	}
}

// 调整堆
func heapify(d []data, root, size int) {
	largest := root     // 初始化最大值为根节点
	left := 2*root + 1  // 左子节点
	right := 2*root + 2 // 右子节点
	if left < size {
		compare++ // 比较次数加一
		if d[left].value > d[largest].value {
			largest = left
		}
	}
	if right < size {
		compare++ // 比较次数加一
		if d[right].value > d[largest].value {
			largest = right
		}
	}
	if largest != root {
		d[root], d[largest] = d[largest], d[root] // 交换根节点和最大值节点
		exchange++                                // 交换次数加一
		heapify(d, largest, size)                 // 递归调整子树
	}
}

// 归并排序函数
func mergeSort(d []data) {
	n := len(d)
	if n <= 1 {
		return
	}

	// 分割数组
	mid := n / 2
	left := make([]data, mid)
	right := make([]data, n-mid)
	copy(left, d[:mid])
	copy(right, d[mid:])

	// 递归排序左右子数组
	mergeSort(left)
	mergeSort(right)

	// 合并左右子数组
	merge(d, left, right)
}

// 合并两个有序数组
func merge(d, left, right []data) {
	i, j, k := 0, 0, 0
	n1, n2 := len(left), len(right)

	for i < n1 && j < n2 {
		compare++
		if left[i].value <= right[j].value {
			d[k] = left[i]
			i++
		} else {
			d[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		d[k] = left[i]
		i++
		k++
	}

	for j < n2 {
		d[k] = right[j]
		j++
		k++
	}

	exchange += n1 + n2 // 交换次数加上已合并的元素数量
}

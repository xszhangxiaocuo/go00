package main

import (
	"bufio"
	"container/list"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

/*
基于哈夫曼编码的文件压缩存储
对一个给定的文本文件，对其进行哈夫曼编码，并计算压缩率。
2023年5月12日17:51:34
*/

type HuffmanTree struct {
	char    rune
	freq    float64
	parents *HuffmanTree
	left    *HuffmanTree
	right   *HuffmanTree
}

func buildTree(freqMap map[rune]int, sum int) *HuffmanTree {
	pq := make([]*HuffmanTree, len(freqMap))
	x := 0
	for key, value := range freqMap {
		pq[x] = &HuffmanTree{key, float64(value) / float64(sum), nil, nil, nil}
		x++
	}
	sort.Slice(pq, func(i, j int) bool {
		return pq[i].freq < pq[j].freq
	})
	l := &list.List{}
	for _, v := range pq {
		l.PushBack(v)
		fmt.Printf("%c:%f\n", v.char, v.freq)
	}
	for l.Len() > 1 {
		a := l.Front().Value.(*HuffmanTree)
		l.Remove(l.Front())
		b := l.Front().Value.(*HuffmanTree)
		l.Remove(l.Front())
		parent := &HuffmanTree{0, a.freq + b.freq, nil, b, a}
		a.parents = parent
		b.parents = parent
		tmp := l.Front()
		for tmp != nil {
			if tmp.Value.(*HuffmanTree).freq >= a.freq+b.freq {
				l.InsertBefore(parent, tmp)
				break
			}
			tmp = tmp.Next()
		}
		if tmp == nil {
			l.PushBack(parent)
		}

	}
	return l.Front().Value.(*HuffmanTree)
}

func generateCodes(root *HuffmanTree, code string, codes *map[rune]string) {
	if root == nil {
		return
	}
	if root.char != 0 {
		(*codes)[root.char] = code
	} else {
		generateCodes(root.left, code+"0", codes)
		generateCodes(root.right, code+"1", codes)
	}
}

func toBinary(code string) []byte {
	result := make([]byte, 0)
	for i := 0; i < len(code); i += 8 {
		end := i + 8
		if end > len(code) {
			end = len(code)
		}
		b, _ := strconv.ParseUint(code[i:end], 2, 8)
		result = append(result, byte(b))
	}
	return result
}

func main() {
	file, _ := ioutil.ReadFile("lab5/huffmandata.txt")
	freqMap := make(map[rune]int)
	for _, char := range string(file) { //统计各字符出现的次数
		freqMap[char]++
	}
	root := buildTree(freqMap, len(file)) //创建哈夫曼树
	codes := make(map[rune]string)
	generateCodes(root, "", &codes) //遍历哈夫曼树得到字母编码
	fmt.Println("哈夫曼编码: ", codes)
	outFile, _ := os.Create("lab5/huffmandata.zipe")
	writer := bufio.NewWriter(outFile)
	for _, char := range string(file) {
		binary.Write(writer, binary.LittleEndian, toBinary(codes[char])) //二进制写入文件
	}
	writer.Flush()
	oldSize := len(file) * 8
	newSize := 0
	for _, char := range string(file) {
		newSize += len(codes[char])
	}
	compressionRatio := float64(newSize) / float64(oldSize) * 100
	fmt.Printf("压缩率: %.2f%%\n", compressionRatio)
}

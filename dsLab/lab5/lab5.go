package main

import (
	"bufio"
	"container/heap"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/*
基于哈夫曼编码的文件压缩存储
对一个给定的文本文件，对其进行哈夫曼编码，并计算压缩率。
2023年5月12日17:51:34
*/

type HuffmanTree struct {
	char  rune
	freq  int
	left  *HuffmanTree
	right *HuffmanTree
}

type PriorityQueue []*HuffmanTree

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanTree))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func buildTree(freqMap map[rune]int) *HuffmanTree {
	pq := make(PriorityQueue, len(freqMap))
	i := 0
	for key, value := range freqMap {
		pq[i] = &HuffmanTree{key, value, nil, nil}
		i++
	}
	heap.Init(&pq)
	for pq.Len() > 1 {
		a := heap.Pop(&pq).(*HuffmanTree)
		b := heap.Pop(&pq).(*HuffmanTree)
		heap.Push(&pq, &HuffmanTree{0, a.freq + b.freq, a, b})
	}
	return heap.Pop(&pq).(*HuffmanTree)
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
	root := buildTree(freqMap) //创建哈夫曼树
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

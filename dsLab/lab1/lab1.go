package main

import (
	"bufio"
	"dsLab/lab1/myMap"
	"fmt"
	"io"
	"math"
	"os"
)

/*
编写程序，计算任意两篇文章的相似度
输入为英文文章，且只根据单词出现频率计算余弦相似度
2023年3月24日10:23:12
*/

func main() {
	str1 := readFile("input1.txt")
	str2 := readFile("input2.txt")

	vector1 := count(str1) //统计各种不同单词出现的次数
	vector2 := count(str2)

	result := cal(vector1, vector2)
	fmt.Printf("%f", result)
}

// 统计单词数
func count(str string) *myMap.MyMap {
	words := myMap.CreatMap() //记录出现过的单词
	word := make([]byte, 0)
	var ch byte
	for i := 0; i < len(str); i++ {
		ch = str[i]
		if ch < 'A' || ch > 'Z' {
			ch -= 32
		}
		if (ch < '0' || ch > '9') && (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') {
			if words.Get(string(word)) != nil { //单词已经存在就更新出现次数
				words.Put(string(word), words.Get(string(word)).(int)+1)
			} else {
				words.Put(string(word), 1)
			}
			word = make([]byte, 0)
			continue
		}
		word = append(word, ch)
	}
	if len(word) != 0 {
		if words.Get(string(word)) != nil { //单词已经存在就更新出现次数
			words.Put(string(word), words.Get(string(word)).(int)+1)
		} else {
			words.Put(string(word), 1)
		}
	}
	return words
}

// 计算相似度
func cal(vector1 *myMap.MyMap, vector2 *myMap.MyMap) float64 {
	var m, num1, num2 int
	var n, n1, n2 float64
	keys := vector1.GetKey() //保存两个map中最大的key的个数
	bigMap := vector1
	smallMap := vector2
	if vector1.Size < vector2.Size { //判断出两个map中key个数较多的那个并更新相应的变量
		keys = vector2.GetKey()
		bigMap, smallMap = smallMap, bigMap
	}
	for _, s := range keys { //遍历key个数较多的map中的所有key
		num1 = bigMap.Get(s).(int)
		if smallMap.Get(s) == nil {
			num2 = 0
		} else {
			num2 = smallMap.Get(s).(int)
		}
		m += num1 * num2
		n1 += float64(num1 * num1)
		n2 += float64(num2 * num2)
	}

	keys = smallMap.GetKey()
	for _, s := range keys { //遍历key个数较多的map中的所有key
		num2 = smallMap.Get(s).(int)
		if bigMap.Get(s) != nil {
			continue
		} else {
			num1 = 0
		}
		m += num1 * num2
		n1 += float64(num1 * num1)
		n2 += float64(num2 * num2)
	}

	n = math.Sqrt(n1) * math.Sqrt(n2)
	return float64(m) / n
}

func readFile(fileName string) string {
	str := ""
	inputFile, inputErr := os.Open("./lab1/" + fileName)
	if inputErr != nil {
		fmt.Println("file open failed!")
		return ""
	}
	defer inputFile.Close()

	read := bufio.NewReader(inputFile)

	for {
		tmp, readErr := read.ReadString('\n')
		if readErr == io.EOF {
			break
		}

		if readErr != nil {
			fmt.Println("read file failed!")
			return ""
		}
		str = str + tmp

	}

	return str
}

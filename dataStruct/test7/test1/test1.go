package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
给定两个非空集合A和B，集合的元素为30000以内的正整数，编写程序求A-B。

输入格式:
输入为三行。第1行为两个整数n和m，分别为集合A和B包含的元素个数，1≤n, m ≤10000。第2行表示集合A，为n个空格间隔的正整数，
每个正整数不超过30000。第3行表示集合B，为m个空格间隔的正整数，每个正整数不超过30000。

输出格式:
输出为一行整数，表示A-B，每个整数后一个空格，各元素按递增顺序输出。若A-B为空集，则输出0，0后无空格。
2023年 3月13日 星期一 14时04分57秒
*/

func main() {
	var n, m int
	fmt.Scan(&n, &m) //输入数组大小
	data1 := make([]int, 0)
	data2 := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	nums := strings.Split(line, " ")
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 32)
		if err == nil {
			data1 = append(data1, int(n))
		}
	}

	line, _ = reader.ReadString('\n')
	line = strings.Trim(line, "\r\n")
	nums = strings.Split(line, " ")
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 32)
		if err == nil {
			data2 = append(data2, int(n))
		}
	}

	result := make([]int, 0)
	a, b := 0, 0
	for {
		if a >= n {
			break
		}
		if b >= m {
			for i := a; i < n; i++ {
				result = append(result, data1[i])
			}
			break
		}

		if data1[a] < data2[b] {
			result = append(result, data1[a])
			a++
		} else if data1[a] == data2[b] {
			a++
			b++
		} else {
			b++
		}

	}

	if len(result) == 0 {
		fmt.Print("0")
		return
	}

	for i, num := range result {
		fmt.Print(num)
		if i != len(result) {
			fmt.Print(" ")
		}
	}

}

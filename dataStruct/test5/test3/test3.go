package main

import "fmt"

/*
已知一组数据，采用顺序存储结构存储，其中所有的元素为整数。设计一个算法，删除元素值在[x,y]之间的所有元素
2023年3月9日16:26:48
*/

func main() {
	var n int
	fmt.Scan(&n) //输入数组大小
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i]) //输入数组元素
	}
	var x, y int
	fmt.Scan(&x, &y) //输入范围
	if x > y {
		return
	}

	index := 0
	for i := 0; i < n; i++ {
		if data[i] <= x || data[i] >= y {
			data[index] = data[i]
			index++
		}
	}
	data = data[:index]
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i])
		if i != len(data)-1 {
			fmt.Print(" ")
		}
	}
}

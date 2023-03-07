package main

import "fmt"

/*
编写函数，用不多于 3n/2 的平均比较次数，在一个顺序表 A 中找出最大和最小值的元素。
2023年 3月 7日 星期二 13时08分17秒
*/

func main() {
	var n int
	fmt.Scan(&n) //输入数组大小
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i]) //输入数组元素
	}

	max := data[0]
	min := data[0]
	for i := 1; i < n; i++ {
		if data[i] > max {
			max = data[i]
		} else if data[i] < min {
			min = data[i]
		}
	}
	fmt.Println(max)
	fmt.Println(min)
}

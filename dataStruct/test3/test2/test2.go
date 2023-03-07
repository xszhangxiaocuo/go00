package main

import "fmt"

/*
编写函数，从一个顺序表A中删除元素值在 x和y(x≤y)之间的所有元素。
2023年 3月 7日 星期二 12时25分42秒
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
		fmt.Println("x的值不能大于y的值")
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
		fmt.Print(data[i], " ")
	}
}

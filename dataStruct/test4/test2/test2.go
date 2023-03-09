package main

import (
	"fmt"
)

/*
以顺序表作存储结构，实现线性表的创建、插入。

输入格式:
输入分三行

第一行 元素个数
第二行 元素的值。元素间以空格分隔。
第三行 待插入的位置和元素值
具体格式参看输入样例

输出格式:
输出分两行

第一行 插入前的线性表
第二行 插入后的线性表。 如因插入位置错误失败，输出Insert position error! 如因为表满插入失败，输出OVERFLOW!

2023年3月8日08:39:56
*/

const MAXSIZE = 10

type elemtype int
type sqList struct {
	elem   []elemtype
	length int
}

func main() {
	var n int
	fmt.Scan(&n) //输入数组大小
	list := new(sqList)
	var data elemtype
	for i := 0; i < n; i++ {
		fmt.Scan(&data) //输入数组元素
		list.elem = append(list.elem, data)
	}
	list.length = n
	var x int
	var y elemtype
	fmt.Scan(&x, &y) //要插入的位置以及元素,x为应插入位置的下标+1

	print(list.elem)

	if x > n+1 || x < 1 {
		fmt.Print("Insert position error!")
		return
	} else if n == MAXSIZE {
		fmt.Print("OVERFLOW!")
		return
	}

	tmp := make([]elemtype, n+1)
	copy(tmp, list.elem[:x-1])
	tmp[x-1] = y
	copy(tmp[x:n+1], list.elem[x-1:n])
	fmt.Print("After:(")
	for i, d := range tmp {
		fmt.Print(d)
		if i == n {
			fmt.Print(")")
		} else {
			fmt.Print(",")
		}
	}

}

func print(data []elemtype) {
	fmt.Print("Before:(")
	for i, d := range data {
		fmt.Print(d)
		if i == len(data)-1 {
			fmt.Print(")\n")
		} else {
			fmt.Print(",")
		}
	}
}

package main

import (
	"fmt"
)

/*
以顺序表作存储结构，实现线性表的创建、删除。

输入格式:
输入分三行

第一行 元素个数
第二行 元素的值。元素间以空格分隔。
第三行 待删除元素位置
具体格式参看输入样例

输出格式:
输出分两行

第一行 删除前的线性表

第二行 删除后的线性表。如因删除位置错误失败，输出Delete position error!。

2023年3月8日09:56:07
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
	fmt.Scan(&x) //x为应插入位置的下标+1

	print(list.elem)

	if x > n || x < 1 {
		fmt.Print("Delete position error!")
		return
	}

	tmp := make([]elemtype, n-1)
	copy(tmp, list.elem[:x-1])
	copy(tmp[x-1:n-1], list.elem[x:n])
	fmt.Print("After:(")
	for i, d := range tmp {
		fmt.Print(d)
		if i == len(tmp)-1 {
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

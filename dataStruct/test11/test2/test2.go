package main

import (
	"container/list"
	"fmt"
)

/*
假设以1和0分别表示入栈和出栈操作。栈的初态和终态均为空，入栈和出栈的操作序列可表示为仅由1和0组成的序列，称可以操作且满足栈终态为空的序列为合法序列
，否则称为非法序列。例如10110100就是一个合法序列，而10010110是一个非法序列，因为第2次出栈时栈已空，该操作无法进行。
请编写程序判断给定的操作序列是否合法。

输入格式:
输入为2行，第1行为一个不超过100的正整数n，表示操作序列中操作的个数，第2行为给定的操作序列，为空格间隔的n个整数，每个整数均为0或1。

输出格式:
输出为一个整数，若输入序列合法，则输出1，若输入序列不合法，则输出0。
2023年4月3日14:33:11
*/

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	stack := &list.List{}
	for _, a := range arr {
		if a == 1 {
			stack.PushBack(1)
		} else {
			if stack.Len() == 0 {
				fmt.Println("0")
				return
			} else {
				stack.Remove(stack.Back())
			}
		}
	}
	if stack.Len() == 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}

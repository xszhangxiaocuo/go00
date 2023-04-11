package main

import (
	"container/list"
	"fmt"
)

/*
有一个死胡同，宽度刚好只能让一辆汽车通过，偏偏老有汽车开到死胡同来，这下麻烦了，最先开来的汽车要最后才能倒退出去。给定一个汽车开来的序列和一个可能
的倒车出去的序列，请判断汽车能否都倒退出去，若能则输出Yes，否则输出No。

输入格式:
首先输入一个整数T，表示测试数据的组数，然后是T组测试数据。每组测试数据首先输入一个正整数n（n≤10），代表开来的汽车数，然后输入2n个整数，其中，
前n个整数表示汽车开来的序列，后n个整数表示汽车可能倒出的序列。

输出格式:
对于每组测试，判断能否倒车出该死胡同，若能则输出“Yes”，否则输出“No”。引号不必输出。
*/

func main() {
	var m, n int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&n)
		enter := make([]int, n)
		leave := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&enter[j])
		}
		for j := 0; j < n; j++ {
			fmt.Scan(&leave[j])
		}
		if canleave(enter, leave) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func canleave(enter, leave []int) bool {
	stack := &list.List{}
	index := 0
	for _, e := range enter {
		stack.PushBack(e)
		//对车的进入顺序进行入栈，每次入栈都将栈顶元素与车子离开顺序的数组leave对比，如果相等就执行出栈操作，直到栈顶元素与当前leave[index]不相等
		for stack.Len() != 0 && stack.Back().Value == leave[index] {
			stack.Remove(stack.Back())
			index++
		}
	}
	if stack.Len() == 0 {
		return true
	}
	return false
}

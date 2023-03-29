package main

import (
	"container/list"
	"fmt"
)

/*
给定一个长度为 N 的整数数列，输出每个数左边第一个比它小的数，如果不存在则输出 −1。

输入格式:
第一行包含整数 N，表示数列长度。

第二行包含 N 个整数，表示整数数列。

输出格式:
共一行，包含 N 个整数，其中第 i 个数表示第 i 个数的左边第一个比它小的数，如果不存在则输出 −1。
2023年3月29日08:43:25
*/

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	tmp := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		arr[i] = tmp
	}
	result := find(arr)
	for _, num := range result {
		fmt.Print(num, " ")
	}
}

func find(nums []int) []int {
	n := len(nums)
	result := make([]int, 0)
	stack := new(list.List)
	stack.PushBack(nums[0])
	result = append(result, -1)
	for i := 1; i < n; i++ {
		peek := stack.Back().Value
		tmp := nums[i]
		flag := false
		if peek.(int) >= tmp { //栈顶元素大于当前元素就移除栈顶元素并将当前较小元素入栈
			for j := len(result) - 1; j >= 0; j-- { //遍历result数组，找到当前元素左边比它小的第一个元素
				if result[j] < tmp && result[j] != -1 {
					result = append(result, result[j])
					flag = true
					break
				}
			}
			stack.Remove(stack.Back())
			stack.PushBack(nums[i])
			if flag {
				continue
			}
			result = append(result, -1)

		} else {
			result = append(result, peek.(int))
			stack.PushBack(tmp)
		}
	}
	return result
}

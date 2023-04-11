package main

import "fmt"

/*
已知有两个等长的非降序序列S1, S2, 设计函数求S1与S2并集的中位数。有序序列A0,A1,⋯,AN−1的中位数指A(N−1)/2的值,即第⌊(N+1)/2⌋个数（A0为第1个数）

输入格式:
输入分三行。第一行给出序列的公共长度N（0<N≤100000），随后每行输入一个序列的信息，即N个非降序排列的整数。数字用空格间隔。

输出格式:
在一行中输出两个输入序列的并集序列的中位数。

2023年4月11日14:34:27
*/

func main() {
	var n int
	fmt.Scan(&n)
	s1 := make([]int, n)
	s2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s1[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&s2[i])
	}
	p1 := 0
	p2 := 0
	result := make([]int, 2*n)
	index := 0
	for p1 < n && p2 < n {
		for p1 < n && p2 < n && s1[p1] < s2[p2] {
			result[index] = s1[p1]
			if index == (2*n-1)/2 {
				fmt.Println(result[(2*n-1)/2])
				return
			}
			index++
			p1++

		}
		for p1 < n && p2 < n && s1[p1] >= s2[p2] {
			result[index] = s2[p2]
			if index == (2*n-1)/2 {
				fmt.Println(result[(2*n-1)/2])
				return
			}
			index++
			p2++
		}
	}
	for p1 < n {
		result[index] = s1[p1]
		index++
		p1++
	}
	for p2 < n {
		result[index] = s2[p2]
		index++
		p2++
	}
	fmt.Println(result[(2*n-1)/2])
}

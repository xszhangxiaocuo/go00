package main

import "fmt"

/*
排列与组合是常用的数学方法，其中组合就是从 n 个元素中抽出 r 个元素(不分顺序且 r≤n)，
我们可以简单地将 n 个元素理解为自然数 1,2,…,n，从中任取 r 个数。
现要求你用递归的方法输出所有组合。

输入格式:
一行两个自然数 n,r(1<n<21，1≤r≤n)。

输出格式:
所有的组合，每一个组合占一行且其中的元素按由小到大的顺序排列，每个元素占三个字符的宽度，右对齐，所有的组合也按字典顺序。

输入样例:

2023年3月29日11:24:22
*/
var n, r int

func main() {
	fmt.Scan(&n, &r)
	find(1, r)
}

func find(start int, length int) {
	if start > r {
		return
	}
	for i := start; i < n; i++ {
		fmt.Print(i, " ")
		for j := 1; j < length; j++ {
			find(start+1, r-j)
		}
	}

}

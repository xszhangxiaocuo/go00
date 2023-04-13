package main

import "fmt"

/*
求整数集合A与整数集合B的交集。

输入格式:
输入有三行：
第一行是A和B的元素个数m和n（m,n <=100）；
第二行是集合A的m个元素；
第三行是集合A的n个元素。

输出格式:
输出交集的所有元素（按照在A集合出现的顺序输出，最后一个输出后面没有空格）。
若交集为空，输出“NULL”。

2023年 4月14日 星期五 00时06分18秒
*/

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	A := make(map[int]bool, m)
	B := make(map[int]bool, n)
	var k int
	keys := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&k)
		A[k] = true
		keys[i] = k
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		B[k] = true
	}
	result := make([]int, 0)
	for _, key := range keys {
		if B[key] {
			result = append(result, key)
		}
	}
	if len(result) == 0 {
		fmt.Println("NULL")
		return
	}
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
		if i != len(result)-1 {
			fmt.Print(" ")
		}
	}
}

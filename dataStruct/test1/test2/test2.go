package main

import "fmt"

/*
我们要求找出具有下列性质数的个数(包含输入的正整数 n)。先输入一个正整数 n,然后对此正整数按照如下方法进行处理：
1.本身不作任何处理；
2.在它的左边加上一个正整数,但该正整数不能超过原数的一半；
3.加上数后,继续按此规则进行处理,直到不能再加正整数为止。
例如输入n=6,
满足条件的数为:
6，16，26，126，36，136

2023年3月4日20:09:06
*/
func main() {
	var n int
	fmt.Scan(&n)
	test(n)
}

func test(n int) {
	result := 1
	var arr [501]int
	arr[1] = 1
	for j := 1; j <= n/2; j++ {
		tmp := 0
		for i := 1; i <= j/2; i++ {
			tmp += arr[i]
		}
		arr[j] = tmp + 1
		result += arr[j]
	}
	fmt.Print(result)
}

//type tree struct {
//	leaves []*tree
//	num    int
//}
//
//func main() {
//	var n int
//	fmt.Scan(&n)
//	test(n)
//}
//
//func test(n int) {
//	result := 1
//	if n == 1 {
//		fmt.Print(result)
//		return
//	}
//	root := new(tree)
//	var count func(int)
//	count = func(num int) {
//		if num < 1 {
//			return
//		}
//		for i := 1; i <= num/2; i++ {
//			tmp := new(tree)
//			tmp.num = i
//			root.leaves = append(root.leaves, tmp)
//			result++
//			count(i)
//		}
//
//	}
//	count(n)
//	fmt.Print(result)
//}

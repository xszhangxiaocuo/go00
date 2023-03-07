package main

import "fmt"

/*
编写函数，将一个顺序表 A（有n个元素且任何元素均不为 0），
分拆成两个顺序表B和C。使A中大于0的元素存放在B中，小于0的元素存放在 C 中，返回顺序表B和C。
2023年 3月 7日 星期二 12时48分42秒
*/

func main() {
	var n int
	fmt.Scan(&n) //输入数组大小
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i]) //输入数组元素
	}

	separate := func() ([]int, []int) {
		var tmpB, tmpC []int
		for i := 0; i < n; i++ {
			if A[i] > 0 {
				tmpB = append(tmpB, A[i])
			} else {
				tmpC = append(tmpC, A[i])
			}
		}
		return tmpB, tmpC
	}

	B, C := separate()
	fmt.Println(B)
	fmt.Println(C)
}

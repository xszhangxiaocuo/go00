package main

import "fmt"

/*
一个数组a中存有n（>0）个整数，将每个整数循环向右移m（≥0）个位置，最后m个数循环移至最前面的m个位置。
第一行给出n和m，分别表示数据的个数和需要右移的个数。其中n的值不超过1000000，m为int型非负整数。
2023年3月9日16:29:35
*/

// 反转一个切片
func reverse(s []int) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

// 循环右移一个切片
func rightShift(s []int, m int) {
	n := len(s)
	m %= n         // 防止m大于n
	reverse(s)     // 反转整个切片
	reverse(s[:m]) // 反转前m个元素
	reverse(s[m:]) // 反转剩余的元素
}

func main() {
	var n, m int // 输入n和m
	fmt.Scan(&n, &m)
	a := make([]int, n)      // 创建长度为n的切片a
	for i := 0; i < n; i++ { // 输入a中的元素
		fmt.Scan(&a[i])
	}

	rightShift(a, m) // 循环右移a

	for i := 0; i < n; i++ { // 输出a中的元素
		if i > 0 {
			fmt.Print(" ") // 元素之间用空格分隔
		}
		fmt.Print(a[i])
	}
}

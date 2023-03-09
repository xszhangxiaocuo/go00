package main

import (
	"fmt"
)

/*
总所周知 LC 是集训队买单最快的人，现在 LC 有 n 张面额从 1 到 10 的钞票，他想请你帮他算算 1 到 10 金额的钞票分别有多少张。

输入格式:
第一行中给出1个正整数 n(1<=n<=10) 代表金库里有 n 张钞票。
第二行给出 n 个正整数，第 i 个整数 Xi 代表第 i 张钞票的金额。

输出格式:
在一行内输出 10 个整数，第 i 个整数代表金额为 i 的钞票的数量，两个数之间用一个空格隔开，注意行末不要有空格
2023年3月6日14:38:20
*/

func main() {
	var n int
	fmt.Scan(&n)

	var tmp int
	count := make(map[int]int, 10)

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		count[tmp]++
	}
	for i := 1; i <= 10; i++ {
		v := count[i]
		fmt.Print(v)
		if i != 10 {
			fmt.Print(" ")
		}
	}
}

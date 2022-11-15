package main

import (
	"fmt"
	"math"
)

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果
如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0
假设环境不允许存储 64 位整数（有符号或无符号）
2022年10月27日20:04:39
*/

func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if !infer(x) {
		return 0
	}
	result := 0

	n := 0
	for x != 0 {
		n = x % 10
		x /= 10
		result = result*10 + n
		if !infer(result) {
			return 0
		}

	}
	return result
}

func infer(x int) bool {
	return x >= int(math.Pow(-2, 31)) && x < int(math.Pow(2, 31))
}

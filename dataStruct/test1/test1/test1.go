package main

import (
	"fmt"
)

/*
对于给定的正整数n（1<=n<=10），求1～n构成的集合的所有子集（幂集）。
2023年3月3日19:00:20
*/

func main() {
	var n uint
	fmt.Scan(&n)
	test(n)
}

func test(n uint) {
	var result [][]int
	var sub []int //用来临时存放一组结果
	result = append(result, sub)
	for i := 1; i < 1<<n; i++ {

		sub = make([]int, 0)
		t := 0
		num := i
		for num != 0 {
			t++
			if num%2 == 1 {
				sub = append(sub, t)
			}
			num /= 2
		}
		result = append(result, sub)
	}

	for _, tmp := range result {
		fmt.Print("{")
		for i, r := range tmp {
			if i != len(tmp)-1 {
				fmt.Print(r, ",")
			} else {
				fmt.Print(r)
			}

		}
		fmt.Println("}")
	}

}

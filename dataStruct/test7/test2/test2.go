package main

import "fmt"

/*
An array of size n ≤ 106 is given to you. There is a sliding window of size k which is moving from the very left of the
array to the very right. You can only see the k numbers in the window. Each time the sliding window moves rightwards by
one position.
2023年 3月13日 星期一 14时47分28秒
*/

func main() {
	var n, k int     //k为窗口大小
	fmt.Scan(&n, &k) //输入数组大小
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i]) //输入数组元素
	}
	max := data[0]
	min := data[0]
	allMax := make([]int, 0)
	allMin := make([]int, 0)
	for i := 0; i < n-k+1; i++ {
		max, min = find(data[i : i+k])
		allMax = append(allMax, max)
		allMin = append(allMin, min)
	}

	for i, num := range allMin {
		fmt.Print(num)
		if i != len(allMin) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	for i, num := range allMax {
		fmt.Print(num)
		if i != len(allMax) {
			fmt.Print(" ")
		}
	}
}

func find(data []int) (int, int) {
	max := data[0]
	min := data[0]
	for _, num := range data {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	return max, min
}

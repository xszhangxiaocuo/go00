package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
给你一个非空列表，返回此列表中 第三大的数 。如果不存在，则返回列表中最大的数。
这本是一道很简单的题目，但是Alan想请你动动脑：设计一个时间复杂度 O(n) 的解决方案。

输入格式:
输入一行数字。
1 <= len(nums) <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
注意，要求返回第三大的数，是指第三大且唯一出现的数。
2023年 3月 9日 星期四 13时44分02
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := make([]int64, 0)
	if scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err == nil {
				arr = append(arr, int64(n))
			}
		}
	}

	fmt.Print(find(arr))
}

func find(nums []int64) int64 {
	var max1, max2, max3 int64 = math.MinInt64, math.MinInt64, math.MinInt64
	t := 0
	for _, num := range nums {
		if num == max1 || num == max2 || num == max3 {
			continue
		}

		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
			t++
		} else if num > max2 {
			max3 = max2
			max2 = num
			t++
		} else if num > max3 {
			max3 = num
		}

	}

	if max3 != math.MinInt64 {
		return max3
	} else {
		return max1
	}
}

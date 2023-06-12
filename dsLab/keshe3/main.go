package main

import "fmt"

/*
0/1背包问题
动态规划
2023年6月9日15:37:39
*/

func main() {
	volumes := []int{1, 8, 4, 3, 5, 2}
	target := 10
	findSubsets(volumes, target)
}

func findSubsets(volumes []int, target int) {
	n := len(volumes)

	// 创建一个二维数组来存储状态
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1) //dp[i][j]表示在前i个物品中是否存在一组物品的总体积等于j
	}

	// 初始化第一行,除了dp[0][0]其它都为false
	dp[0][0] = true

	// 动态规划的状态转移
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= volumes[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-volumes[i-1]]
			}
		}
	}
	for _, bools := range dp {
		fmt.Println(bools)
	}
	// 回溯找出所有解
	subset := []int{}
	backtrack(dp, volumes, target, n, subset)
}

func backtrack(dp [][]bool, volumes []int, target int, i int, subset []int) {
	if i == 0 && target == 0 {
		printSubset(subset)
		return
	}

	// 不选择第i个物品
	if i > 0 && dp[i-1][target] { //前i-1物品中已经有组合满足总体积为target
		newSubset := make([]int, len(subset))
		copy(newSubset, subset)
		backtrack(dp, volumes, target, i-1, newSubset)
	}

	// 选择第i个物品
	if i > 0 && target-volumes[i-1] >= 0 && dp[i-1][target-volumes[i-1]] {
		newSubset := make([]int, len(subset))
		copy(newSubset, subset)
		newSubset = append(newSubset, volumes[i-1])
		backtrack(dp, volumes, target-volumes[i-1], i-1, newSubset)
	}
}

func printSubset(subset []int) {
	fmt.Print("(")
	for i, val := range subset {
		fmt.Print(val)
		if i < len(subset)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println(")")
}

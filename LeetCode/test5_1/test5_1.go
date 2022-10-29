package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
使用动态规划
2022年10月27日18:24:34
*/

func main() {
	s := "abab"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	//空串或长度为1，直接返回s
	if s == "" || len(s) == 1 {
		return s
	}
	strLen := len(s)
	maxLen := 0                  //最大回文串长度
	maxLeft := 0                 //最大回文串左边界下标
	maxRight := 0                //最大回文串右边界下标
	dp := make([][]bool, strLen) //dp[l][r]表示下标从l到到r的字符串是否是回文串
	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
	}
	//长度为1的字符串都是回文串
	for i := 0; i < strLen; i++ {
		dp[i][i] = true
	}
	//从r=1开始循环，因为一定有r>l
	for r := 1; r < strLen; r++ {
		for l := 0; l < r; l++ {
			//当前左边界的字符等于右边界的字符,并且从l+1到r-1的字符串为回文串时（特殊情况为当前的子串长度小于等于3）,才有dp[l][r]=true
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true

				if r-l+1 > maxLen {
					maxLen = r - l + 1
					maxLeft = l
					maxRight = r
				}
			}
		}
	}
	return s[maxLeft : maxRight+1]

}

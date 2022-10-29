package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
使用中心扩散
2022年10月26日21:55:01
*/

func main() {
	s := "abab"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	strLen := len(s)
	left := 0    //左扩散的边界下标
	right := 0   //右扩散的边界下标
	length := 0  //当前回文串的长度
	maxLen := 0  //最大回文串长度
	maxLeft := 0 //最大回文串的左边界

	for i := 0; i < strLen; i++ {
		left = i - 1
		right = i + 1
		//由中心开始左扩散，如果遇到与中心字符不相等的就退出
		for left >= 0 && (s[left] == s[i]) {
			left--
		}
		//由中心开始右扩散，如果遇到与中心字符不相等的就退出
		for right < strLen && (s[right] == s[i]) {
			right++
		}
		//由当前的left和right所在位置分别同时左右扩散，如果在扩散过程中两个字符不相等就退出
		for left >= 0 && right < strLen && (s[left] == s[right]) {
			left--
			right++
		}
		//最后退出时的下标肯定是不符合条件的，所以要往前复原一次
		left++
		right--
		length = right - left + 1
		//如果当前len为最大值，就覆盖maxLen和maxLeft的值
		if length > maxLen {
			maxLen = length
			maxLeft = left
		}
	}

	return s[maxLeft : maxLeft+maxLen]
}

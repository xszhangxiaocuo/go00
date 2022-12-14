package main

import "fmt"

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列，之后，你的输出需要从左往右逐行读取，产生出一个新的字符串。
用二维数组模拟
2022年10月27日18:24:45
*/

func main() {
	s := "PAYPALISHIRING"
	n := 3
	fmt.Println(convert(s, n))
}

func convert(s string, numRows int) string {
	length := len(s)
	//如果行数为1或者字符串长度小于等于行数，直接返回
	if numRows == 1 || length <= numRows {
		return s
	}
	sortStr := make([][]byte, numRows)
	t := 2*numRows - 2
	col := (length/t + 1) * (numRows - 1) //多加上一个周期，避免有些情况下最后一个周期不完整导致数组列数不够
	//创建numRows行，col列的二维数组
	for i := range sortStr {
		sortStr[i] = make([]byte, col)
	}
	x, y := 0, 0
	for i, ch := range s {
		sortStr[x][y] = byte(ch)
		if i%t < numRows-1 { //在前半个周期内，行数一直增加
			x++
		} else { //后半个周期内，一直向右上方移动，直到x=0
			x--
			y++
		}
	}
	var result []byte
	for _, arr := range sortStr {
		for _, ch := range arr {
			if ch > 0 {
				result = append(result, ch)
			}
		}
	}
	return string(result)
}

package main

import (
	"bytes"
	"fmt"
)

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列，之后，你的输出需要从左往右逐行读取，产生出一个新的字符串。
用二维数组模拟(优化：压缩空白数组，节省空间)
2022年10月29日11:29:34
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
	x := 0
	for i, ch := range s {
		sortStr[x] = append(sortStr[x], byte(ch))
		if i%t < numRows-1 { //在前半个周期内，行数一直增加
			x++
		} else { //后半个周期内，一直向右上方移动，直到x=0
			x--
		}
	}

	return string(bytes.Join(sortStr, nil)) //将二维切片sortStr连接起来组成一个新的一维切片，并以sep作为分隔符，这里nil表示没有分隔符
}

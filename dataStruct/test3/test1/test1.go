package main

import "fmt"

/*
已知顺序表的元素按非降序排列 。请编写算法，删除表中的重复元素。例如，原表为（1，1，2，3，3，3，4，5，5），
经算法处理后，表为（1，2，3，4，5）。要求算法的空间复杂度为O(1)，不需输出表元素的值。顺序表描述如下：
typedef struct t
{  int data[MAXLEN];  int len;
}LIST;
2023年 3月 7日 星期二 11时40分11秒
*/

const MAXLEN = 100

type list struct {
	data [MAXLEN]int
	len  int
}

func main() {
	l := new(list)
	var n int
	fmt.Scan(&n)
	l.len = n
	for i := 0; i < n; i++ {
		fmt.Scan(&l.data[i])
	}

	tmp := l.data[0]
	index := 1
	for i := 1; i < n; i++ {
		if l.data[i] != tmp {
			l.data[index] = l.data[i]
			tmp = l.data[i]
			index++
		}
	}
	l.len = index
	for i := 0; i < l.len; i++ {
		fmt.Print(l.data[i], " ")
	}
}

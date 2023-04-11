package main

import "fmt"

/*
本题要求计算单链表倒数n个结点的乘积。例如，给出单链表1 2 3 4 5，则倒数2个结点的乘积为20。

输入格式:
输入有2行，第一个行为2个非负整数m和n。其中m为链表结点个数，n为链表倒数结点的数量。题目保证计算结果在int范围内。
第二行为链表的m个数，以空格分隔。

输出格式:
在一行中输出倒数n个结点的乘积。
*/

type List struct {
	data int
	next *List
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	if m == 0 || n == 0 {
		fmt.Println("0")
		return
	}
	head := creat(m)
	t := 0
	result := 1
	for head.next != nil {
		if t >= m-n {
			result *= head.next.data
		}
		t++
		head = head.next
	}
	fmt.Println(result)
}

// 有头节点的尾插法创建链表
func creat(n int) *List {
	head := new(List)
	tail := head
	for i := 0; i < n; i++ {
		tmp := new(List)
		fmt.Scan(&tmp.data)
		tail.next = tmp
		tail = tmp
	}
	return head
}

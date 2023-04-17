package main

import "fmt"

/*
编写程序，删除单链表head中的重复结点。

2023年4月17日12:39:29
*/

type List struct {
	data int
	next *List
}

func main() {
	var n int
	fmt.Scan(&n)
	head := creat(n)
	del(head)
	print(head)
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

// 输出链表
func print(head *List) {
	for head.next != nil {
		fmt.Print(head.next.data, " ")
		head = head.next
	}
	fmt.Println()
}

func del(head *List) {
	flag := make([]bool, 1000)
	for head.next != nil {
		if flag[head.next.data] {
			head.next = head.next.next
		} else {
			flag[head.next.data] = true
			head = head.next
		}
	}
}

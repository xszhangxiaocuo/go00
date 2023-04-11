package main

import "fmt"

/*
编写程序，删除元素递增排列的链表head中值大于min且小于max的所有元素
2023年4月10日17:46:45
*/

type List struct {
	data int
	next *List
}

func main() {
	var n int
	var min, max int
	fmt.Scan(&n)
	head := creat(n)
	fmt.Scan(&min, &max)
	del(head, min, max)
	print(head)
}

func del(head *List, min int, max int) {
	list := head
	minp := head
	maxp := head
	minflag := false
	for list.next != nil {
		if list.next.data > min {
			minflag = true
		} else if !minflag {
			minp = list.next
		}
		if list.next.data < max {
			maxp = list.next
		}
		list = list.next
	}
	minp.next = maxp.next
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

package main

import (
	"fmt"
)

/*
编写程序，在头指针为head的单链表中，把值为key的结点插入到值为a的结点之前，若不存在a，就把key插入到表尾。
2023年4月11日12:15:48
*/

type List struct {
	data int
	next *List
}

func main() {
	var n, key, a int
	fmt.Scan(&n)
	head := creat(n)
	fmt.Scan(&key, &a)
	insert(head, key, a)
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

func insert(head *List, key int, a int) {
	for head.next != nil {
		if head.next.data == a {
			break
		}
		head = head.next
	}
	tmp := &List{
		data: key,
		next: nil,
	}
	if head.next == nil {
		head.next = tmp
	} else {
		tmp.next = head.next
		head.next = tmp
	}
}

package main

import "fmt"

/*
编写程序，将任意一个带头结点的单链表原地逆序。
2023年3月28日14:38:04
*/

type List struct {
	num  int
	next *List
}

func main() {
	var n int
	fmt.Scan(&n)
	list := creat(n)
	reverse(list)
	printList(list)
}

// 有头节点的尾插法创建链表
func creat(n int) *List {
	head := new(List)
	tail := head

	for i := 0; i < n; i++ {
		tmp := new(List)
		fmt.Scan(&tmp.num)
		tail.next = tmp
		tail = tmp
	}
	return head
}

// 打印链表
func printList(head *List) {
	list := head
	for list.next != nil {
		fmt.Print(list.next.num)
		if list.next.next != nil {
			fmt.Print(" ")
		}
		list = list.next
	}
	fmt.Println()
}

// 逆序链表
func reverse(head *List) *List {
	list := head.next
	var pre *List //保存前一个节点
	for list != nil {
		next := list.next //保存下一个节点
		list.next = pre   //将当前节点的next指向前一个节点
		pre = list        //将pre更新为当前节点
		list = next       //将list更新为下一个节点
	}
	head.next = pre
	return head
}

package main

import "fmt"

/*
已知一个单链表是一个有头节点的链表，编写一个函数将该单链表复制到另一个单链表中
2023年 3月18日 星期六 21时25分13秒
*/

type elemtype byte

type List struct {
	key  elemtype
	next *List
}

func main() {
	n := 10
	list := creat(n)
	result := listCopy(list)
	print(list)
	print(result)
}

// 有头节点的尾插法创建链表
func creat(n int) *List {
	head := new(List)
	tail := head

	for i := 0; i < n; i++ {
		tmp := new(List)
		tmp.key = elemtype(i + 48)
		tail.next = tmp
		tail = tmp
	}

	return head
}

// 输出链表
func print(head *List) {
	for head.next != nil {
		fmt.Printf("%c ", head.next.key)
		head = head.next
	}
	fmt.Println()
}

// 复制链表（遍历源链表，取出源链表中每个节点的key重新创建新的节点加入到目标链表中，而不是直接复制源链表节点的地址）
func listCopy(list *List) *List {
	head := new(List)
	tail := head
	for list.next != nil {
		tmp := new(List)
		tmp.key = list.next.key
		tail.next = tmp
		tail = tmp
		list = list.next
	}
	return head
}

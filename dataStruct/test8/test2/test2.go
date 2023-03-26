package main

import "fmt"

/*
输入若干个不超过100的整数，建立单链表，然后将链表中所有结点的链接方向逆置，要求仍利用原表的存储空间。输出逆置后的单链表。

输入格式:
首先输入一个整数T，表示测试数据的组数，然后是T组测试数据。每组测试数据在一行上输入数据个数n及n个不超过100的整数。

输出格式:
对于每组测试，输出逆置后的单链表，每两个数据之间留一个空格。
2023年 3月22日 星期三 15时57分40秒
*/

type List struct {
	num  int
	next *List
}

func main() {
	var m, n int
	fmt.Scan(&m)
	heads := make([]*List, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&n)
		heads[i] = creat(n)
	}
	for i := 0; i < m; i++ {
		heads[i] = reverse(heads[i])
		printList(heads[i])
	}

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

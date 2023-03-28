package main

import "fmt"

/*
编写程序，求两个递增有序的单链表A与单链表B的差集。已知每个单链表中均不存在重复元素。差集：集合的差A-B，指包含所有属于集合A而不属于集合B的元素。
2023年3月28日14:49:24
*/

type List struct {
	num  int
	next *List
}

func main() {
	var n int
	fmt.Scan(&n)
	list1 := creat(n) //集合A
	fmt.Scan(&n)
	list2 := creat(n) //集合B
	printList(del(list1, list2))
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

// 求集合A-B
func del(head1 *List, head2 *List) *List {
	list1 := head1
	list2 := head2

	for list1.next != nil && list2.next != nil {
		if list1.next.num == list2.next.num { //找到相同的节点然后将该节点从链表A中断开
			list1.next = list1.next.next
			list2 = list2.next
		} else if list1.next.num < list2.next.num { //当前list1.next小于list2.next时，后移list1指针
			list1 = list1.next
		}
	}

	return head1
}

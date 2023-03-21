package main

import (
	"fmt"
)

/*
编写程序，对两个一元多项式求和，并输出“和多项式”。
默认按幂的降序输入
2023年 3月18日 星期六 21时41分45秒
*/

type List struct {
	base int //系数
	pow  int //幂次
	next *List
}

func main() {
	var m, n int
	fmt.Scan(&m)
	head1 := creat(m)
	fmt.Scan(&n)
	head2 := creat(n)
	fmt.Println("第一个多项式为：")
	printList(head1)
	fmt.Println("第二个多项式为：")
	printList(head2)
	resultHead := new(List)
	resultTail := resultHead
	n1, n2 := head1, head2 //分别是第一个链表和第二个链表的指针
	for n1.next != nil && n2.next != nil {
		if n1.next.pow == n2.next.pow { //两个节点幂次相等就相加
			if n1.next.base+n2.next.base != 0 { //两个元素相加不为0
				resultTail = add(resultTail, n1.next.base+n2.next.base, n1.next.pow)
			}
			n1 = n1.next
			n2 = n2.next
		} else if n1.next.pow > n2.next.pow { //n1.next指向的幂次高于n2.next说明该高次幂无需合并，直接加入result即可
			resultTail = add(resultTail, n1.next.base, n1.next.pow)
			n1 = n1.next
		} else if n1.next.pow < n2.next.pow { //n1.next指向的幂次低于n2.next说明该高次幂无需合并，直接加入result即可
			resultTail = add(resultTail, n2.next.base, n2.next.pow)
			n2 = n2.next
		}
	}
	for n1.next != nil {
		add(resultTail, n1.next.base, n1.next.pow)
		n1 = n1.next
	}
	for n2.next != nil {
		add(resultTail, n2.next.base, n2.next.pow)
		n2 = n2.next
	}
	fmt.Println("两个多项式的和为：")
	printList(resultHead)
}

// 输入两个多项式并用链表存储
func creat(n int) *List {
	var base, pow int
	head := new(List)
	tail := head

	for i := 0; i < n; i++ {
		fmt.Scan(&base, &pow)
		tmp := &List{base, pow, nil}
		tail.next = tmp
		tail = tmp
	}

	return head
}

// 给list添加节点
func add(tail *List, base int, pow int) *List {
	tmp := &List{base, pow, nil}
	tail.next = tmp
	tail = tmp
	return tail
}

// 输出链表
func printList(head *List) {
	t := 0
	for head.next != nil {
		if t != 0 { //判断是否是第一个元素
			if head.next.base > 0 { //大于0输出+号，负数会自动输出
				fmt.Print("+")
			}

		}
		t++
		if head.next.pow == 0 { //x的幂次为0时，直接输出系数
			fmt.Print(head.next.base)
		} else if head.next.base != 0 { //系数为0不输出
			fmt.Print(head.next.base, "x^", head.next.pow)
		}
		head = head.next
	}
	fmt.Println()
}

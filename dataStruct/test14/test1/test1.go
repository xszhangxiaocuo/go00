package main

import "fmt"

/*
已知两个非降序链表序列S1与S2，设计函数构造出S1与S2的交集新链表S3。

输入格式:
输入分两行，分别在每行给出由若干个正整数构成的非降序序列，用−1表示序列的结尾（−1不属于这个序列）。数字用空格间隔。

输出格式:
在一行中输出两个输入序列的交集序列，数字间用空格分开，结尾不能有多余空格；若新链表为空，输出NULL。
2023年 4月13日 星期四 23时30分58秒
*/

type List struct {
	data int
	next *List
}

func main() {
	list1 := creat()
	list2 := creat()

	result := del(list1, list2)
	print(result)
}

func del(head1 *List, head2 *List) *List {
	head := &List{}
	tail := head

	for head1.next != nil && head2.next != nil {
		for head1.next != nil && head2.next != nil && head1.next.data < head2.next.data {
			head1 = head1.next
		}
		for head1.next != nil && head2.next != nil && head1.next.data > head2.next.data {
			head2 = head2.next
		}
		if head1.next != nil && head2.next != nil && head1.next.data == head2.next.data {
			tmp := &List{
				data: head1.next.data,
				next: nil,
			}
			tail.next = tmp
			tail = tmp
			head1 = head1.next
			head2 = head2.next
		}
	}

	for head1.next != nil {
		if head1.next.data == head2.data {
			tmp := &List{
				data: head1.next.data,
				next: nil,
			}
			tail.next = tmp
			tail = tmp
			break
		}
		head1 = head1.next
	}

	for head2.next != nil {
		if head1.data == head2.next.data {
			tmp := &List{
				data: head1.next.data,
				next: nil,
			}
			tail.next = tmp
			tail = tmp
			break
		}
		head2 = head2.next
	}

	return head
}

func creat() *List {
	head := new(List)
	tail := head
	var data int
	for {
		fmt.Scan(&data)
		if data == -1 {
			break
		}
		tmp := new(List)
		tmp.data = data
		tail.next = tmp
		tail = tmp
	}

	return head
}

// 输出链表
func print(head *List) {
	if head.next == nil {
		fmt.Println("NULL")
		return
	}
	for head.next != nil {
		fmt.Print(head.next.data)
		head = head.next
		if head.next != nil {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

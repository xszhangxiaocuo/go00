package main

import (
	"fmt"
	"math"
)

/*
编写一个完整的程序，实现将两个非递减集合LA和LB合并生成具有相同属性的集合LC。
要求：
1）写出数据的定义
2）写出每个函数的具体实现
3）写出主函数

输入格式:
输入在一行中给出LA的元素个数，再依次输入值添加到LA中；
接着输入LB的元素个数，再依次输入值添加到LB中。
注意添加的值，要依次递增，否则提示错误，要求重新输入该值。

输出格式:
LA/LB中的元素添加完成后，显示LA/LB线性表；
最后输出生成的LC。
2023年 3月22日 星期三 16时52分18秒
*/

type List struct {
	num  int
	next *List
}

func main() {
	var n int
	fmt.Scan(&n)
	ok1, la := creat(n)
	for !ok1 {
		ok1, la = creat(n)
	}
	fmt.Scan(&n)
	ok2, lb := creat(n)
	for !ok2 {
		ok2, la = creat(n)
	}

	lc := merge(la, lb)

	printList(la)
	printList(lb)
	printList(lc)
}

// 有头节点的尾插法创建链表
func creat(n int) (bool, *List) {
	head := new(List)
	tail := head
	pre := math.MinInt64
	for i := 0; i < n; i++ {
		tmp := new(List)
		fmt.Scan(&tmp.num)
		if pre > tmp.num {
			return false, nil
		}
		tail.next = tmp
		tail = tmp
	}
	return true, head
}

// 打印链表
func printList(head *List) {
	list := head
	for list.next != nil {
		fmt.Print(list.next.num)
		fmt.Print(" ")
		list = list.next
	}
	fmt.Println()
}

func merge(la *List, lb *List) *List {
	head := new(List)
	tail := head

	tmp := new(List)
	for la.next != nil && lb.next != nil {
		tmp = new(List)
		if la.next.num <= lb.next.num {
			tmp.num = la.next.num
			la = la.next
		} else {
			tmp.num = lb.next.num
			lb = lb.next
		}
		tail.next = tmp
		tail = tmp
	}

	for la.next != nil {
		tmp = new(List)
		tmp.num = la.next.num
		la = la.next
		tail.next = tmp
		tail = tmp
	}
	for lb.next != nil {
		tmp = new(List)
		tmp.num = lb.next.num
		lb = lb.next
		tail.next = tmp
		tail = tmp
	}

	return head
}

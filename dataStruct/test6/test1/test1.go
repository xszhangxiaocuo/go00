package main

import "fmt"

/*
单链表基础操作：构造一个带头结点的单链表。其每个结点中记录着一个字符型的键值key(设键值唯一)。编写函数，完成表1的操作。
表1 线性表的基本操作
⑴ 初始化线性表

⑵ 输出线性表

⑶ 取表中的第i个元素的键值

⑷ 从表中删除指定位置的元素

⑸ 从表中删除指定键值的元素

⑹ 向表的头部添加键值为key的元素

⑺ 向表的尾部添加键值为key的元素

⑻ 向表中指定的位置pos处添加键值为key的元素

⑼ 在表中中搜索键值为key的元素，看其是否存在

2023年 3月18日 星期六 20时19分55
*/

type elemtype byte

type List struct {
	key  elemtype
	next *List
}

func main() {
	n := 10
	list := creat(n)
	print(list)
	fmt.Printf("第%d个元素的键值为：%c\n", n, get(list, n))
	delIndex(list, n-2)
	print(list)
	delKey(list, '5')
	print(list)
	addHead(list, 'A')
	print(list)
	addTail(list, 'B')
	print(list)
	addPos(list, 1, 'C')
	print(list)
	fmt.Print("D:", find(list, 'D'))
	fmt.Print("\nA:", find(list, 'A'))
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

// 取出链表第n个元素的键值
func get(head *List, n int) elemtype {
	var key elemtype
	t := 0

	for head.next != nil {
		t++
		if t == n {
			key = head.next.key
			break
		}
		head = head.next
	}

	return key
}

// 从表中删除指定位置的元素,index为下标[0,n)
func delIndex(head *List, index int) {
	t := 0
	for head.next != nil {
		if t == index {
			head.next = head.next.next
			break
		}
		t++
		head = head.next
	}
}

// 从表中删除指定键值的元素
func delKey(head *List, key elemtype) {
	list := head
	for list.next != nil {
		if list.next.key == key {
			list.next = list.next.next
			break
		}
		list = list.next
	}
}

// 向表的头部添加键值为key的元素
func addHead(head *List, key elemtype) {
	tmp := &List{
		key:  key,
		next: head.next,
	}
	head.next = tmp
}

// 向表的尾部添加键值为key的元素
func addTail(head *List, key elemtype) {
	for head.next != nil {
		head = head.next
	}
	head.next = &List{
		key:  key,
		next: head.next,
	}
}

// 向表中指定的位置pos[0,n)处添加键值为key的元素
func addPos(head *List, pos int, key elemtype) {
	t := 0
	for head.next != nil {
		if t == pos {
			break
		}
		t++
		head = head.next
	}
	head.next = &List{
		key:  key,
		next: head.next,
	}
}

// 在表中中搜索键值为key的元素，看其是否存在
func find(head *List, key elemtype) bool {
	for head.next != nil {
		if head.next.key == key {
			return true
		}
		head = head.next
	}
	return false
}

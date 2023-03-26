package main

import "fmt"

/*
17世纪的法国数学家加斯帕在《数目的游戏问题》中也讲了这样一个故事：15个教徒和15个非教徒在深海上遇险，必须将一半的人投入海中，
其余的人才能幸免于难，于是想了一个办法：30个人围成一圆圈，从第一个人开始依次报数，每数到第九个人就将他扔入大海，如此循环进行直到仅余15个人为止。
【问题】怎样的安排才能使每次投入大海的都是非教徒？请编程解决这一n(1≤n≤30)个人的跳海问题。要求分别用两种线性表的存储结构来解决。
采用循环链表实现
2023年 3月21日 星期二 15时13分25秒
*/

type Person struct {
	id   int //id从0开始编号
	next *Person
}

var STEP = 8 //计数值为8时丢掉一个人
var N = 30

func main() {
	head := creat(N)
	printList(head)
	drop(head)
	printList(head)

}

// 有头节点的尾插法创建循环链表
func creat(n int) *Person {
	head := new(Person)
	tail := head

	for i := 0; i < n; i++ {
		tmp := new(Person)
		tmp.id = i
		tail.next = tmp
		tail = tmp
	}
	tail.next = head
	return head
}

// 打印循环链表
func printList(head *Person) {
	list := head
	for list.next != head {
		fmt.Printf("%d ", list.next.id+1)
		list = list.next
	}
	fmt.Println()
}

// 从链表中删掉被丢进海中的id，留下的id即为所求答案
func drop(head *Person) {
	list := head

	for i := 0; i < N/2; i++ { //丢掉一半的人
		for j := 0; j < STEP; j++ { //每次往后走STEP步,此时的list.next即是下一个要被丢掉的人
			if list.next == head { //头节点无意义，需要再往后走一步
				list = list.next
			}
			list = list.next
		}
		list.next = list.next.next
	}
}

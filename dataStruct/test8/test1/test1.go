package main

import "fmt"

/*
编号为1…N的N个小朋友玩游戏，他们按编号顺时针围成一圈，按顺时针次序报数，从第1个人报到第M个人出列；然后再从下个人开始报到第M+1个人出列；
再从下一个人开始报到第M+2个人出列……以此类推不断循环，直至最后一人出列。请编写程序按顺序输出出列人的编号。
输入格式:
输入为2个整数，分别表示N、M（1≤N,M,K≤10000）。

输出格式:
输出为一行整数，为出列人的编号。每个整数后一个空格。
2023年 3月22日 星期三 12时49分26秒
*/

type Person struct {
	id   int //id从0开始编号
	next *Person
}

var STEP = 0 //计数值为8时丢掉一个人
var N = 0

func main() {
	fmt.Scan(&N, &STEP)
	STEP--
	head := creat(N)
	drop(head)
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

// 从链表中删掉被丢进海中的id，留下的id即为所求答案
func drop(head *Person) {
	list := head

	for i := 0; i < N; i++ {
		for j := 0; j < STEP; j++ { //每次往后走STEP步,此时的list.next即是下一个要被丢掉的人
			if list.next == head { //头节点无意义，需要再往后走一步
				list = list.next
			}
			list = list.next
			if list.next == head { //头节点无意义，需要再往后走一步
				list = list.next
			}
		}
		fmt.Print(list.next.id+1, " ")
		list.next = list.next.next
		STEP++
	}
}

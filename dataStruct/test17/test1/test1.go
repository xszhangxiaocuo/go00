package main

import "fmt"

/*
编写程序，对任一单链表，计算链表中数据只出现一次的结点个数，要求空间复杂度为O(1).
2023年5月13日21:20:41
*/

type Node struct {
	data int
	next *Node
}

func countNode(head *Node) int {
	count := 0
	current := head

	for current != nil { //遍历链表
		flag := false //标记当前节点有没有出现过
		runner := head

		for runner != current { //用另一个指针遍历链表寻找有无重复出现的节点
			if runner.data == current.data {
				flag = true
				break
			}
			runner = runner.next
		}

		if !flag { //没有重复出现就加1
			count++
		}

		current = current.next
	}

	return count
}

func main() {
	// 创建链表：1 -> 2 -> 4 -> 2 -> 4 -> 5 -> 1
	head := &Node{data: 1}
	head.next = &Node{data: 2}
	head.next.next = &Node{data: 4}
	head.next.next.next = &Node{data: 2}
	head.next.next.next.next = &Node{data: 4}
	head.next.next.next.next.next = &Node{data: 5}
	head.next.next.next.next.next.next = &Node{data: 1}

	count := countNode(head)
	fmt.Println("没有重复出现的节点个数:", count)
}

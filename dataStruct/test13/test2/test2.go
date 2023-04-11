package main

import (
	"container/list"
	"fmt"
)

/*
两端分别是一条入口（Entrance）轨道和一条出口（Exit）轨道，它们之间有N条平行的轨道。每趟列车从入口可以选择任意一条轨道进入，最后从出口离开。
在图中有9趟列车，在入口处按照{8，4，2，5，3，9，1，6，7}的顺序排队等待进入。
如果要求它们必须按序号递减的顺序从出口离开，则至少需要多少条平行铁轨用于调度？

输入格式：
输入第一行给出一个整数N (2 ≤ N ≤105)，下一行给出从1到N的整数序号的一个重排列。数字间以空格分隔。

输出格式：
在一行中输出可以将输入的列车按序号递减的顺序调离所需要的最少的铁轨条数。
2023年4月11日15:01:04
*/

func main() {
	var n int
	fmt.Scan(&n)
	train := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&train[i])
	}
	var queues []*list.List
	queue := &list.List{}
	queues = append(queues, queue)
	currentTrain := n
	flag := false
	for i := 0; i < n; i++ {
		if queue.Len() == 0 || queue.Back().Value.(int) > train[i] {
			queue.PushBack(train[i])
		} else {
			f := true
			ff := true
			if len(queues) == 1 {
				ff = false
			}
			for _, q := range queues {
				if q.Len() == 0 && ff {
					queue = q
					queue.PushBack(train[i])
					f = false
					break
				}
			}
			if f {
				queue = &list.List{}
				queue.PushBack(train[i])
				queues = append(queues, queue)
			}
		}
		if train[i] == currentTrain {
			queue.Remove(queue.Front())
			currentTrain--
			for !flag {
				tmp := currentTrain
				for _, q := range queues {
					for q.Len() != 0 && q.Front().Value.(int) == currentTrain {
						q.Remove(q.Front())
						currentTrain--
					}
				}
				if tmp == currentTrain {
					flag = true
				}
			}
			flag = false
		}
	}
	fmt.Println(len(queues))

}

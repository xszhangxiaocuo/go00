package main

import (
	"dsLab/lab3/Queue"
	"fmt"
)

/*
操作系统的时间片轮转调度算法
2023年4月7日10:36:39
*/

const slices = 5 //时间片长度5ms

func main() {
	q := createQueue()
	result := Queue.NewQueue()
	var time uint = 0
	for !q.IsEmpty() {
		tmp := q.Pop()
		if tmp.T == 0 {
			tmp.Start = time
		}
		time += slices
		tmp.T++
		if tmp.T*slices < tmp.Run {
			q.PushNode(tmp)
		} else {
			tmp.Finish = time
			tmp.Ti = time
			tmp.Wi = float64(time) / float64(tmp.Run)
			result.PushNode(tmp)
		}
	}
	fmt.Println("success")
	for !result.IsEmpty() {
		tmp := result.Pop()
		fmt.Println(tmp.Value, " ", tmp.Arrive, " ", tmp.Run, " ", tmp.Start, " ", tmp.Finish, " ", tmp.Ti, " ", tmp.Wi)
	}
}

func createQueue() *Queue.Queue {
	q := Queue.NewQueue()
	q.Push("A", 0, 20)
	q.Push("B", 0, 10)
	q.Push("C", 0, 15)
	q.Push("D", 0, 5)
	return q
}

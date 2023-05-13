package main

import (
	"container/list"
	"fmt"
)

/*
编写程序，实现对以下有向图的邻接表存储，并输出存储的图（自定义输出形式）。给出从任意出发的BFS遍历序列。
2023年5月13日22:23:53
*/

type Node struct {
	value  int
	visted bool //标记该节点是否被访问过
	nodes  []*Node
}

func main() {
	var n int
	fmt.Println("请输入顶点个数：")
	fmt.Scan(&n)
	nodes := make([]Node, n)
	for i := 0; i < n; i++ { //输入顶点
		fmt.Scan(&nodes[i].value)
	}
	fmt.Print("请输入边数：")
	fmt.Scan(&n)
	for i := 0; i < n; i++ { //输入边的信息
		a := 0
		b := 0
		fmt.Print("请输入第", i+1, "条边的顶点对：")
		fmt.Scan(&a, &b)
		nodes[a].nodes = append(nodes[a].nodes, &nodes[b])
	}
	fmt.Println()
	//输出存储的图
	for i := 0; i < len(nodes); i++ {
		fmt.Print("顶点", i, "邻接顶点有：")
		for _, node := range nodes[i].nodes {
			fmt.Print(node.value, " ")
		}
		fmt.Println()
	}

	fmt.Print("bfs:")
	bfs(&nodes[0])
}

func bfs(head *Node) {
	queue := &list.List{}
	queue.PushFront(head)
	head.visted = true
	fmt.Print(head.value, " ")
	for queue.Len() > 0 { //队列为空说明遍历结束
		tmp := queue.Front()        //取出队首元素
		queue.Remove(queue.Front()) //移除队首元素
		for _, node := range tmp.Value.(*Node).nodes {
			if node.visted { //该节点被访问过了直接下一个
				continue
			}
			node.visted = true
			queue.PushFront(node) //没访问过的节点入列
			fmt.Print(node.value, " ")
		}
	}
}

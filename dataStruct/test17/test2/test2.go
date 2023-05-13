package main

import "fmt"

/*
编写程序，实现对以下无向图的邻接矩阵存储，并输出存储的图（以下左图为一种输出图的形式）。给出从任意出发的DFS遍历序列。
2023年5月13日21:49:25
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
		//无向图两个顶点互通
		nodes[a].nodes = append(nodes[a].nodes, &nodes[b])
		nodes[b].nodes = append(nodes[b].nodes, &nodes[a])
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

	fmt.Print("dfs:")
	dfs(&nodes[0])
}

func dfs(head *Node) {
	stak := make([]*Node, 0)
	stak = append(stak, head) //起点入栈
	head.visted = true
	for len(stak) > 0 { //栈为空说明遍历结束
		tmp := stak[len(stak)-1]  //取出栈顶元素
		stak = stak[:len(stak)-1] //弹出栈顶元素
		fmt.Print(tmp.value, " ")
		for _, node := range tmp.nodes {
			if node.visted { //该节点被访问过了直接下一个
				continue
			}
			node.visted = true
			stak = append(stak, node) //没访问过的节点入栈
		}
	}
}

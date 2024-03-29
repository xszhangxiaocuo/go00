package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
	"math/rand"
	"time"
)

/*
旅行推销员问题，找到一条经过所有节点的最短路径。本题的解法是基于贪婪算法实现的近似算法，在处理大规模数据时可能不是最优解
2023年6月13日23:10:23
*/

const Inf = math.MaxInt32 //Inf表示两点不可到达

type Node struct {
	X, Y float64
}

type Edge struct {
	Start, End int
	Weight     int
}

type Graph struct {
	Vertices int
	Nodes    []Node
	Edges    []Edge
}

func main() {
	vertices := 5
	graph := NewGraph(vertices)
	//graph.AddEdge(0, 1, 2)
	//graph.AddEdge(0, 2, 4)
	//graph.AddEdge(1, 2, 1)
	//graph.AddEdge(1, 3, 3)
	//graph.AddEdge(1, 4, 7)
	//graph.AddEdge(2, 3, 5)
	//graph.AddEdge(3, 4, 6)
	//graph.AddEdge(2, 5, 2)
	//graph.AddEdge(3, 6, 1)
	//graph.AddEdge(4, 7, 5)
	//graph.AddEdge(5, 6, 6)
	//graph.AddEdge(6, 7, 4)
	//graph.AddEdge(5, 8, 3)
	//graph.AddEdge(6, 9, 2)
	//graph.AddEdge(8, 9, 1)

	//graph.AddEdge(0, 1, 2)
	//graph.AddEdge(0, 2, 5)
	//graph.AddEdge(1, 2, 1)
	//graph.AddEdge(1, 3, 4)
	//graph.AddEdge(1, 4, 7)
	//graph.AddEdge(2, 3, 5)
	//graph.AddEdge(3, 4, 6)

	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 3)
	graph.AddEdge(1, 4, 7)
	graph.AddEdge(2, 3, 5)
	graph.AddEdge(3, 4, 6)

	path, distance := FindPath(graph)

	fmt.Println("最短路径：")
	for _, v := range path {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Printf("最短路径距离：%d\n", distance)

	DrawMap(graph, path)
}

func NewGraph(vertices int) *Graph {
	rand.Seed(time.Now().UnixNano()) //生成随机坐标用于绘制地图
	nodes := make([]Node, vertices)
	for i := 0; i < vertices; i++ {
		nodes[i] = Node{
			X: rand.Float64() * 800,
			Y: rand.Float64() * 600,
		}
	}

	return &Graph{
		Vertices: vertices,
		Nodes:    nodes,
		Edges:    make([]Edge, 0),
	}
}

func (g *Graph) AddEdge(src, dest, weight int) {
	edge := Edge{
		Start:  src,
		End:    dest,
		Weight: weight,
	}
	g.Edges = append(g.Edges, edge)
}

func FindPath(graph *Graph) ([]int, int) {
	start := 0                              //从编号为0的节点出发
	visited := make([]bool, graph.Vertices) //标记节点是否被访问过
	path := make([]int, 0)                  //记录最短路径
	path = append(path, start)
	visited[start] = true
	flag := false //每次寻路都有两种方案，false表示直接寻找距离当前节点最近的节点，true表示返回上一个节点另寻一条路
	for !Finish(visited) {
		flag = false
		curr := path[len(path)-1] //当前节点
		next := -1                //下一个节点
		minDist := Inf            //当前节点到附近的最近距离

		// 寻找距离当前节点最近且未访问过的节点
		for v := 0; v < graph.Vertices; v++ {
			if !visited[v] && graph.GetEdgeWeight(curr, v) < minDist {
				next = v
				minDist = graph.GetEdgeWeight(curr, v)
			}
		}
		if len(path)-2 > 0 { //退回到上一个节点重新寻找是否存在一条路径比当前minDist还要小
			minDist += graph.GetEdgeWeight(curr, path[len(path)-2]) //退回到上一个节点要增加一段距离
			roadlen := 0
			roadlen += graph.GetEdgeWeight(curr, path[len(path)-2]) * 2
			curr = path[len(path)-2] //退回上一个节点
			for v := 0; v < graph.Vertices; v++ {
				if !visited[v] && graph.GetEdgeWeight(curr, v)+roadlen < minDist {
					next = v
					minDist = graph.GetEdgeWeight(curr, v) + roadlen
					flag = true //方案二找到了一条比方案一还短的路径就采用方案二
				}
			}
		}

		if next == -1 { //没有找到下一个可以走的节点
			break
		}

		if flag {
			path = append(path, curr)
		}

		// 添加下一个节点到路径中，并标记为已访问
		path = append(path, next)
		visited[next] = true
	}

	return path, CalculatePathDistance(graph, path)
}

// Finish 判断寻路是否结束，如果还有节点为false则没有结束
func Finish(visited []bool) bool {
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

// GetEdgeWeight 获取两点之间的距离，如果两点之间不可到达，返回最大值Inf
func (g *Graph) GetEdgeWeight(src, dest int) int {
	for _, edge := range g.Edges {
		if (edge.Start == src && edge.End == dest) || (edge.Start == dest && edge.End == src) { //无向图，没有起点和终点
			return edge.Weight
		}
	}
	return Inf
}

func CalculatePathDistance(graph *Graph, path []int) int {
	distance := 0
	for i := 0; i < len(path)-1; i++ {
		distance += graph.GetEdgeWeight(path[i], path[i+1])
	}
	return distance
}

func DrawMap(graph *Graph, path []int) {
	const (
		width  = 800
		height = 600
		radius = 10
		margin = 50
	)

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 设置背景为白色
	dc.Clear()

	// 查找最小和最大的坐标值
	minX, minY, maxX, maxY := math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	for _, node := range graph.Nodes {
		if node.X < minX {
			minX = node.X
		}
		if node.X > maxX {
			maxX = node.X
		}
		if node.Y < minY {
			minY = node.Y
		}
		if node.Y > maxY {
			maxY = node.Y
		}
	}

	// 计算缩放比例
	scaleX := (width - 2*margin) / (maxX - minX)
	scaleY := (height - 2*margin) / (maxY - minY)

	// 绘制节点之间的路径
	dc.SetRGB(0, 0, 0) // 设置线条颜色为黑色
	dc.SetLineWidth(1)
	for _, edge := range graph.Edges {
		// 获取节点坐标
		x1 := graph.Nodes[edge.Start].X
		y1 := graph.Nodes[edge.Start].Y
		x2 := graph.Nodes[edge.End].X
		y2 := graph.Nodes[edge.End].Y

		// 缩放坐标到绘图范围内
		scaledX1 := margin + (x1-minX)*scaleX
		scaledY1 := margin + (y1-minY)*scaleY
		scaledX2 := margin + (x2-minX)*scaleX
		scaledY2 := margin + (y2-minY)*scaleY

		dc.DrawLine(scaledX1, scaledY1, scaledX2, scaledY2)
		dc.Stroke()

		// 绘制边的权重
		weightX := (scaledX1 + scaledX2) / 2
		weightY := (scaledY1 + scaledY2) / 2
		dc.DrawStringAnchored(fmt.Sprintf("%d", edge.Weight), weightX, weightY, 0.5, 0.5)
	}

	// 绘制最短路径
	dc.SetRGB(0, 1, 0) // 设置线条颜色为绿色
	dc.SetLineWidth(2)
	for i := 0; i < len(path)-1; i++ {
		v1 := path[i]
		v2 := path[i+1]

		// 获取节点坐标
		x1 := graph.Nodes[v1].X
		y1 := graph.Nodes[v1].Y
		x2 := graph.Nodes[v2].X
		y2 := graph.Nodes[v2].Y

		// 缩放坐标到绘图范围内
		scaledX1 := margin + (x1-minX)*scaleX
		scaledY1 := margin + (y1-minY)*scaleY
		scaledX2 := margin + (x2-minX)*scaleX
		scaledY2 := margin + (y2-minY)*scaleY

		dc.DrawLine(scaledX1, scaledY1, scaledX2, scaledY2)
		dc.Stroke()
	}

	// 绘制节点
	for v, node := range graph.Nodes {
		x := node.X
		y := node.Y

		scaledX := margin + (x-minX)*scaleX
		scaledY := margin + (y-minY)*scaleY

		dc.DrawCircle(scaledX, scaledY, radius)
		dc.SetRGB(0, 0, 0) // 设置节点颜色为黑色
		dc.Fill()
		dc.SetRGB(1, 0, 0) // 设置节点颜色为黑色
		// 绘制节点编号
		dc.DrawStringAnchored(fmt.Sprintf("%d", v), scaledX, scaledY, 0.5, 0.5)
	}

	dc.SavePNG("map.png")
}

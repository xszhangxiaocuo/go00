package main

import "fmt"

/*
请将以下稀疏点阵信息用三元组表进行存储，并：
	*	*	*	*	*	*	*
*								*
							*	*
						*
					*
				*
			*
		*
	*
*
*	*	*	*	*	*	*	*	*	*

(1) 用稀疏矩阵快速转置法对该矩阵进行转置。转置前后的三元组表均以行序为主序。
(2) 以阵列形式输出转置前后的稀疏矩阵，

x y
0 1
0 2
0 3
0 4
0 5
0 6
0 7
1 0
1 8
2 7
2 8
3 6
4 5
5 4
6 3
7 2
8 1
9 0
10 0
10 1
10 2
10 3
10 4
10 5
10 6
10 7
10 8
10 9



2023年4月17日16:32:01
*/

type Point struct {
	x     int
	y     int
	value int
}

func main() {
	points := initPoint(28)
	var src [11][10]int
	var result [10][11]int
	for _, point := range points {
		tmpx := point.x
		tmpy := point.y
		tmpv := point.value
		src[tmpx][tmpy] = tmpv
		result[tmpy][tmpx] = tmpv
	}

	for _, p := range src {
		for _, v := range p {
			if v == 1 {
				fmt.Print("* ")
			} else {
				print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	for _, p := range result {
		for _, v := range p {
			if v == 1 {
				fmt.Print("* ")
			} else {
				print("  ")
			}
		}
		fmt.Println()
	}
}

func initPoint(n int) []Point {
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		point := Point{value: 1}
		fmt.Scan(&point.x, &point.y)
		points[i] = point
	}
	return points
}

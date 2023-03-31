package main

import "fmt"

/*
地图四染色问题
2023年3月31日10:20:22
*/

const citNUm = 7   //城市个数
const colorNum = 4 //颜色个数

var prs = [citNUm][citNUm]int{}
var colors = make([]int, citNUm)
var paint = [citNUm][colorNum]bool{}

func main() {
	initMap()
	dye()
	fmt.Println(colors)
}

// 初始化邻接矩阵
func initMap() {
	i := 0
	arr := [citNUm]int{0, 1, 1, 1, 1, 1, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{1, 0, 0, 0, 0, 1, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{1, 0, 0, 1, 1, 0, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{1, 0, 1, 0, 1, 1, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{1, 0, 1, 1, 0, 1, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{1, 1, 0, 1, 1, 0, 0}
	prs[i] = arr
	i++

	arr = [citNUm]int{0, 0, 0, 0, 0, 0, 0}
	prs[i] = arr
	i++
}

// 地图染色
func dye() {
	pr := 0
	color := 1
	colors[pr] = color //初始化先将第一个省份的颜色涂上1号色,0表示该省份还没有上色
	pr++
	for pr < citNUm {
		for color = 1; color <= colorNum; color++ {
			if canPaint(&pr, color) {
				break
			}
		}
		if color <= colorNum { //当前省份有颜色可以涂
			colors[pr] = color
			pr++
		} else {
			for !back(&pr, color) {

			}
			pr++
		}

	}
}

// 回退
func back(pr *int, color int) bool {
	colors[*pr] = 0
	paint[*pr] = [colorNum]bool{}
	*pr--
	colors[*pr]++
	color = colors[*pr]
	if canPaint(pr, color) { //如果当前颜色可行，就返回true表示结束回退，否则继续往前回退
		return true
	}
	return false
}

// 判断当前省份能否用color着色
func canPaint(pr *int, color int) bool {
	if color > colorNum || paint[*pr][color-1] { //当前颜色已被记录不可用直接退出
		return false
	}
	curPr := 0
	for curPr = 0; curPr < citNUm; curPr++ {
		if prs[*pr][curPr] == 1 && colors[curPr] == color { //两个省份相邻并且已经染过当前的颜色
			paint[*pr][color-1] = true //表示pr城市不能用当前颜色
			return false
		}
	}
	return true //说明当前的color没有被用过，有颜色可以涂
}

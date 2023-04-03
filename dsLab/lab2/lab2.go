package main

import "fmt"

/*
地图四染色问题
2023年3月31日10:20:22
*/

//var provinceName = [34]string{
//	"北京", "天津", "上海", "重庆",
//	"河北", "河南", "云南", "辽宁", "黑龙江",
//	"湖南", "安徽", "山东", "新疆", "江苏",
//	"浙江", "江西", "湖北", "广西", "甘肃",
//	"山西", "内蒙古", "陕西", "吉林", "福建", "贵州",
//	"广东", "青海", "西藏", "四川", "宁夏",
//	"海南", "台湾", "香港", "澳门",
//}
//
//const (
//	//shanxis是陕西，shanxi是山西
//
//	beijing = iota
//	tianjin
//	shanghai
//	chongqing
//	hebei
//	henan
//	yunnan
//	liaoning
//	heilongjiang
//	hunan
//	anhui
//	shandong
//	xinjiang
//	jiangsu
//	zhejiang
//	jiangxi
//	hubei
//	guangxi
//	gansu
//	shanxi
//	neimenggu
//	shanxis
//	jilin
//	fujian
//	guizhou
//	guangdong
//	qinghai
//	xizang
//	sichuan
//	ningxia
//	hainan
//	taiwan
//	xianggang
//	aomen
//)

const cityNum = 34 //城市个数
const colorNum = 3 //颜色个数

var prs = [cityNum][cityNum]int{}
var colors = make([]int, cityNum)

// var colors = [7]int{}
var paint = [cityNum][colorNum]bool{}

func main() {
	initMap()
	dye()
	fmt.Println(colors)
	//colors= [cityNum]int{1, 2, 1, 1, 1, 1, 1}
	fmt.Println("染色结果：", isTrue())
}

// 初始化邻接矩阵
func initMap() {
	//prs = [cityNum][cityNum]int{
	//	{0, 1, 1, 1, 1, 1, 0},
	//	{1, 0, 0, 0, 0, 1, 0},
	//	{1, 0, 0, 1, 1, 0, 0},
	//	{1, 0, 1, 0, 1, 1, 0},
	//	{1, 0, 1, 1, 0, 1, 0},
	//	{1, 1, 0, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0},
	//}

	prs = [cityNum][cityNum]int{
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	//prs[beijing][tianjin] = 1
	//prs[beijing][hebei] = 1
	//
	//prs[tianjin][beijing] = 1
	//prs[tianjin][hebei] = 1
	//
	//prs[shanghai][jiangsu] = 1
	//prs[shanghai][zhejiang] = 1
	//
	//prs[chongqing][sichuan] = 1
	//prs[chongqing][shanxis] = 1
	//prs[chongqing][hubei] = 1
	//prs[chongqing][hunan] = 1
	//prs[chongqing][guizhou] = 1
	//
	//prs[hebei][beijing] = 1
	//prs[hebei][tianjin] = 1
	//prs[hebei][liaoning] = 1
	//prs[hebei][neimenggu] = 1
	//prs[hebei][shanxi] = 1
	//prs[hebei][henan] = 1
	//prs[hebei][shandong] = 1
	//
	//prs[henan][shanxi] = 1
	//prs[henan][hebei] = 1
	//prs[henan][shandong] = 1
	//prs[henan][anhui] = 1
	//prs[henan][hubei] = 1
	//prs[henan][shanxis] = 1
	//
	//prs[yunnan][guizhou] = 1
	//prs[yunnan][guangxi] = 1
	//prs[yunnan][sichuan] = 1
	//prs[yunnan][xizang] = 1
	//
	//prs[liaoning][jilin] = 1
	//prs[liaoning][neimenggu] = 1
	//prs[liaoning][hebei] = 1
	//
	//prs[heilongjiang][jilin] = 1
	//prs[heilongjiang][neimenggu] = 1
	//
	//prs[hunan][hubei] = 1
	//prs[hunan][jiangxi] = 1
	//prs[hunan][guangdong] = 1
	//prs[hunan][guizhou] = 1
	//prs[hunan][guangxi] = 1
	//prs[hunan][chongqing] = 1
	//
	//prs[anhui][henan] = 1
	//prs[anhui][shandong] = 1
	//prs[anhui][jiangsu] = 1
	//prs[anhui][zhejiang] = 1
	//prs[anhui][jiangxi] = 1
	//prs[anhui][hubei] = 1
	//
	//prs[shandong][hebei] = 1
	//prs[shandong][jiangsu] = 1
	//prs[shandong][anhui] = 1
	//prs[shandong][henan] = 1
	//
	//prs[xinjiang][gansu] = 1
	//prs[xinjiang][qinghai] = 1
	//prs[xinjiang][xizang] = 1
	//
	//prs[jiangsu][shanghai] = 1
	//prs[jiangsu][zhejiang] = 1
	//prs[jiangsu][anhui] = 1
	//prs[jiangsu][shandong] = 1
	//
	//prs[zhejiang][shanghai] = 1
	//prs[zhejiang][jiangsu] = 1
	//prs[zhejiang][anhui] = 1
	//prs[zhejiang][jiangxi] = 1
	//prs[zhejiang][fujian] = 1
	//
	//prs[jiangxi][anhui] = 1
	//prs[jiangxi][zhejiang] = 1
	//prs[jiangxi][fujian] = 1
	//prs[jiangxi][guangdong] = 1
	//prs[jiangxi][hunan] = 1
	//prs[jiangxi][hubei] = 1
	//
	//prs[hubei][henan] = 1
	//prs[hubei][anhui] = 1
	//prs[hubei][jiangxi] = 1
	//prs[hubei][hunan] = 1
	//prs[hubei][chongqing] = 1
	//prs[hubei][shanxis] = 1
	//
	//prs[guangxi][guizhou] = 1
	//prs[guangxi][hunan] = 1
	//prs[guangxi][guangdong] = 1
	//prs[guangxi][yunnan] = 1
	//
	//prs[gansu][ningxia] = 1
	//prs[gansu][xinjiang] = 1
	//prs[gansu][shanxis] = 1
	//prs[gansu][sichuan] = 1
	//prs[gansu][qinghai] = 1
	//prs[gansu][neimenggu] = 1
	//
	//prs[shanxi][neimenggu] = 1
	//prs[shanxi][hebei] = 1
	//prs[shanxi][henan] = 1
	//prs[shanxi][shanxis] = 1
	//
	//prs[neimenggu][liaoning] = 1
	//prs[neimenggu][jilin] = 1
	//prs[neimenggu][heilongjiang] = 1
	//prs[neimenggu][hebei] = 1
	//prs[neimenggu][shanxi] = 1
	//prs[neimenggu][shanxis] = 1
	//prs[neimenggu][ningxia] = 1
	//prs[neimenggu][gansu] = 1
	//
	//prs[shanxis][neimenggu] = 1
	//prs[shanxis][shanxi] = 1
	//prs[shanxis][henan] = 1
	//prs[shanxis][hubei] = 1
	//prs[shanxis][chongqing] = 1
	//prs[shanxis][sichuan] = 1
	//prs[shanxis][gansu] = 1
	//prs[shanxis][ningxia] = 1
	//
	//prs[jilin][heilongjiang] = 1
	//prs[jilin][liaoning] = 1
	//prs[jilin][neimenggu] = 1
	//
	//prs[fujian][zhejiang] = 1
	//prs[fujian][jiangxi] = 1
	//prs[fujian][guangdong] = 1
	//
	//prs[guizhou][chongqing] = 1
	//prs[guizhou][hunan] = 1
	//prs[guizhou][guangxi] = 1
	//prs[guizhou][yunnan] = 1
	//prs[guizhou][sichuan] = 1
	//
	//prs[guangdong][hunan] = 1
	//prs[guangdong][jiangxi] = 1
	//prs[guangdong][fujian] = 1
	//prs[guangdong][guangxi] = 1
	//prs[guangdong][aomen] = 1
	//prs[guangdong][xianggang] = 1
	//
	//prs[qinghai][gansu] = 1
	//prs[qinghai][sichuan] = 1
	//prs[qinghai][xizang] = 1
	//prs[qinghai][xinjiang] = 1
	//
	//prs[xizang][xinjiang] = 1
	//prs[xizang][qinghai] = 1
	//prs[xizang][sichuan] = 1
	//prs[xizang][yunnan] = 1
	//
	//prs[sichuan][qinghai] = 1
	//prs[sichuan][gansu] = 1
	//prs[sichuan][shanxis] = 1
	//prs[sichuan][chongqing] = 1
	//prs[sichuan][guizhou] = 1
	//prs[sichuan][yunnan] = 1
	//prs[sichuan][xizang] = 1
	//
	//prs[ningxia][neimenggu] = 1
	//prs[ningxia][shanxis] = 1
	//prs[ningxia][gansu] = 1
	//
	//prs[xianggang][guangdong] = 1
	//prs[aomen][guangdong] = 1
}

// 地图染色
func dye() {
	pr := 0
	color := 1
	colors[pr] = color //初始化先将第一个省份的颜色涂上1号色,0表示该省份还没有上色
	pr++
	for pr < cityNum {
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
	for curPr = 0; curPr < cityNum; curPr++ {
		if prs[*pr][curPr] == 1 && colors[curPr] == color { //两个省份相邻并且已经染过当前的颜色
			paint[*pr][color-1] = true //表示pr城市不能用当前颜色
			return false
		}
	}
	return true //说明当前的color没有被用过，有颜色可以涂
}

func isTrue() bool {

	for i, pr := range prs {
		for j, p := range pr {
			if p == 1 && colors[i] == colors[j] {
				return false
			}
		}
	}

	return true
}

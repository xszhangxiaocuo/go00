package main

import "fmt"

/*
如果跳到了方块上，但没有跳到方块的中心则获得1分；
跳到方块中心时，若上一次的得分为1分或这是本局游戏的第一次跳跃则此次得分为2分，否则此次得分比上一次得分多两分（即连续跳到方块中心时，总得分将+2，+4，+6，+8...）。
1表示此次跳跃跳到了方块上但是没有跳到中心，
2表示此次跳跃跳到了方块上并且跳到了方块中心，
0表示此次跳跃没有跳到方块上（此时游戏结束）。
2023年3月6日13:58:20
*/

func main() {
	t := 0     //连续跳跃到中心的次数
	score := 0 //总分
	for {
		var tmp int
		fmt.Scan(&tmp)
		if tmp == 0 {
			break
		} else if tmp == 1 {
			t = 0
			score++
		} else if tmp == 2 {
			t++
			score += 2 * t
		}
	}
	fmt.Print(score)
}

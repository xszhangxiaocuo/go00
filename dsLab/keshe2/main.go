package main

import (
	"container/list"
	"fmt"
)

/*
农夫过河问题
一个农夫带着—只狼、一只羊和—棵白菜，身处河的南岸。他要把这些东西全部运到北岸。他
面前只有一条小船，船只能容下他和—件物品。只有农夫才能撑船。如果农夫在场，则狼不能吃羊，
羊不能吃白菜，否则狼会吃羊，羊会吃白菜，所以农夫不能留下羊和白菜自己离开，也不能留下狼
和羊自己离开，而狼不吃白菜。
请编写程序：求出农夫将所有的东西运过河的方案。

2023年6月7日15:58:07
*/

const (
	FARMER  int8 = 8
	WOLF    int8 = 4
	CABBAGE int8 = 2
	SHEEP   int8 = 1
)

func main() {
	bfs()
}

func bfs() {
	var status int8 = 0       //0000从高到低分别表示农夫，狼，白菜，羊在南岸
	route := make([]int8, 16) //记录已经考虑过的状态路径
	for i := 0; i < 16; i++ { //初始化为-1
		route[i] = -1
	}
	queue := &list.List{}
	queue.PushBack(status)
	for queue.Len() != 0 && route[15] == -1 { //route
		current := queue.Front().Value.(int8)
		queue.Remove(queue.Front())
		var mover int8
		for mover = 1; mover <= 8; mover <<= 1 { //依次移动农夫，狼，白菜，羊
			if (current&FARMER == 0) == (current&mover == 0) { //农夫和当前要移动的角色在同一侧
				tmp := current ^ (FARMER | mover)
				if (route[tmp] == -1) && isLegal(tmp) { //判断当前状态是否考虑过并且是否合法
					route[tmp] = current
					queue.PushBack(tmp)
				}
			}
		}
	}
	fmt.Println(route)
	if route[15] != -1 {
		var i int8
		for i = 15; i >= 0; i = route[i] {
			fmt.Printf("current location:%4b\n", i)
			if i == 0 {
				break
			}
		}
	} else {
		fmt.Println("no path")
	}
}

func isLegal(status int8) bool {
	flag := false
	//true表示在南岸，false表示在北岸
	farmer := (status & FARMER) == 0
	wolf := (status & WOLF) == 0
	cabbage := (status & CABBAGE) == 0
	sheep := (status & SHEEP) == 0

	if wolf != sheep && cabbage != sheep { //狼和羊不在一起，羊和白菜不在一起
		flag = true
	} else if wolf == sheep { //农夫和狼和羊在一起
		if farmer == wolf {
			flag = true
		}
	} else if sheep == cabbage { //农夫和羊和白菜在一起
		if farmer == sheep {
			flag = true
		}
	}

	return flag
}

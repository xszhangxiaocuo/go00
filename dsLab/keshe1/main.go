package main

import (
	"fmt"
	"math"
	"sort"
)

/*
一元稀疏多项式计算器
⑴ 输入并建立多项式；
⑵ 输出多项式，输出形式为整数序列：n,c1,e1, c2,e2,,,,,,, cn,en,其中 n 是多项式的项数，ci,ei,分别是第 i 项的系数和指数，序列按指数降序排序；
⑶ 实现多项式 a 和 b 相加，建立多项式 a+b；
⑷ 实现多项式 a 和 b 相减，建立多项式 a-b；
⑸ 计算多项式在 x 处的值。
⑹ 计算器的仿真界面。

2023年6月7日14:37:31
*/

type Item struct {
	c float64 //系数
	e int     //指数
}

func main() {
	a := []Item{{1, 1}, {1, 2}, {2, 3}}
	b := []Item{{-1, 1}, {2, 2}, {3, 3}}
	r1 := count(a, b, '+')
	r2 := count(a, b, '-')
	printItems(r1)
	fmt.Println(cal(r1, 1))
	printItems(r2)
	fmt.Println(cal(r2, 1))
}

func count(a []Item, b []Item, cal byte) []Item {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	result := make([]Item, 0)
	m := make(map[int][]Item, 0)

	for _, v := range a {
		m[v.e] = append(m[v.e], v)
	}
	for _, v := range b {
		switch cal {
		case '+':
			m[v.e] = append(m[v.e], v)
		case '-':
			v.c = -v.c
			m[v.e] = append(m[v.e], v)
		}
	}

	for _, v := range m {
		var sumc float64 = 0
		for _, i := range v {
			sumc += i.c
		}
		if sumc != 0 {
			result = append(result, Item{c: sumc, e: v[0].e})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].e > result[j].e
	})
	return result
}

// 计算多项式在x处的值
func cal(items []Item, x float64) float64 {
	if len(items) == 0 {
		return 0
	}
	var sum float64
	for _, v := range items {
		sum += v.c * math.Pow(x, float64(v.e))
	}
	return sum
}

func printItems(items []Item) {
	fmt.Print(len(items), ",")
	for _, i := range items {
		fmt.Print(i.c)
		if i.e != 0 {
			fmt.Print(" ", i.e)
		}
		fmt.Print(",")
	}
	fmt.Println()
}

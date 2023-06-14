package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
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
	var input string
	fmt.Print("请输入多项式A：")
	fmt.Scan(&input)
	a := extractItems(input)
	fmt.Print("请输入多项式B：")
	fmt.Scan(&input)
	b := extractItems(input)

	r1 := count(a, b, '+')
	r2 := count(a, b, '-')
	fmt.Print("A+B=")
	printItems(r1)
	fmt.Print("A-B=")
	printItems(r2)

	var x float64
	fmt.Print("输入x的赋值：")
	fmt.Scan(&x)
	fmt.Printf("当x=%f时，多项式A的值为：%f", x, cal(r1, x))
	fmt.Printf("当x=%f时，多项式B的值为：%f", x, cal(r2, x))
}

func extractItems(str string) []Item {
	var items []Item

	// 去除空格
	str = strings.ReplaceAll(str, " ", "")

	// 匹配每一项,输入的每一项格式必须为cx^e
	re := regexp.MustCompile(`([-+]?\d+\.?\d*)x\^(\d+)`) //([-+]?\d+\.?\d*) 匹配系数部分，可以是正负号、整数或小数。x\^ 匹配"x^",(\d+) 匹配指数部分，一个或多个数字
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		coefficient, _ := strconv.ParseFloat(match[1], 64) //匹配系数
		exponent, _ := strconv.Atoi(match[2])              //匹配指数
		items = append(items, Item{coefficient, exponent})
	}

	return items
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
	for n, i := range items {
		if i.c == 0 {
			continue
		}
		if n != 0 && i.c > 0 {
			fmt.Print("+")
		}
		fmt.Print(i.c)
		if i.e != 0 {
			fmt.Printf("x^%d", i.e)
		}
	}
	fmt.Println()
}

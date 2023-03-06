package main

import "fmt"

/*
小明和小东是生活在二进制世界的人，一天他们分别获得了一个十进制数，他们想通过将该十进制数分解成二进制数，
然后取出其二进制下最小非零数值的部分进行比较，如果小明的更小那么输出"xiaoming"，如果小东的更小输入“xiaodong”，如果一样小输出“same”。
保证所有输入数据 1≤x,y≤1e18.
2023年3月6日14:03:20
*/

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		test(a, b)
	}
}

func test(a int, b int) {
	if a == b {
		fmt.Println("same")
		return
	}

	for a != 0 && b != 0 {
		tmpa := a % 2
		tmpb := b % 2
		if tmpa > tmpb {
			fmt.Println("xiaoming")
			return
		} else if tmpa < tmpb {
			fmt.Println("xiaodong")
			return
		} else if tmpa == tmpb && tmpa == 1 {
			fmt.Println("same")
			return
		}
		a = a >> 1
		b = b >> 1
	}
}

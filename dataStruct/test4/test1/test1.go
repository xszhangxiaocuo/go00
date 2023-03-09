package main

import "fmt"

/*
已知顺序表L递增有序，将X插入到线性表的适当位置上，保证线性表有序。

输入格式:
第1行输入顺序表长度，第2行输入递增有序的顺序表，第3行输入要插入的数据元素X。

输出格式:
对每一组输入，在一行中输出插入X后的递增的顺序表。
2023年3月7日22:47:32
*/

func main() {
	var n int
	fmt.Scan(&n) //输入数组大小
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i]) //输入数组元素
	}
	var x int
	fmt.Scan(&x) //要插入的元素

	left := 0
	right := n - 1
	var m int
	var k = 1
	if x < 0 {
		k = -1
	}
	for left <= right {
		m = (left + right) / 2
		if data[m] >= x { //当前元素大于等于x往前找
			if m > 0 && m < n && data[m-1] <= x {
				break
			}
			right = m - 1
		} else { //当前元素小于x往后找
			if m >= 0 && m < n-1 && data[m+1] > x {
				break
			}
			left = m + 1
		}
	}

	tmp := make([]int, n+1)
	if m == 0 {
		k = 0
	}
	copy(tmp, data[:m+k])
	tmp[m+k] = x
	copy(tmp[m+k+1:n+1], data[m+k:n])
	for _, num := range tmp {
		fmt.Print(num, ",")
	}

}

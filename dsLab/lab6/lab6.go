package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

/*
找出图中入度最多的十个节点
2023年5月13日10:53:34
*/

const MAX = 1894

func main() {
	// 打开 CSV 文件
	f, err := os.Open("lab6/High-School_data_2013.csv")
	if err != nil {
		log.Fatalf("文件打开失败！: %s\n", err.Error())
	}
	defer f.Close()

	records := [MAX][MAX]bool{}
	// 读取 CSV 数据
	r := csv.NewReader(f)
	result := make([]int, MAX)
	resultid := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		resultid[i] = i
	}
	flag := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return
		}
		if flag {
			flag = false
			continue
		}
		x, err := strconv.Atoi(record[0])
		y, err := strconv.Atoi(record[1])
		x--
		y--
		records[x][y] = true
		result[y]++
	}

	selectSort(result, resultid)
	for i := 0; i < 10; i++ {
		fmt.Println(resultid[i]+1, ":", result[i])
	}
}

func selectSort(arr []int, id []int) {
	for i := 0; i < len(arr); i++ {
		k := i
		for j := i + 1; j < len(arr); j++ {
			if arr[k] < arr[j] {
				k = j
			}
		}
		if k != i {
			arr[k], arr[i] = arr[i], arr[k]
			id[k], id[i] = id[i], id[k]
		}
	}
}

package main

import (
	scmap "dsLab/lab8/myMap/SCMap"
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

type telephoneInfo struct {
	id       int
	name     string
	address  string
	phonenum string
}

func main() {
	sc := readInfo()
	for {
		key := ""
		fmt.Println("请输入要查询的用户名：")
		fmt.Scan(&key)
		result := sc.Get(key)
		if result != nil {
			fmt.Println(result)
		} else {
			fmt.Println("用户不存在！")
		}
		fmt.Print("是否继续查询(y/n):")
		fmt.Scan(&key)
		fmt.Println(key)
		if key != "y" && key != "Y" {
			break
		}
	}

}

// 以用户名为关键字建立散列表
func readInfo() *scmap.SCHashTable {
	sc := scmap.CreatSCHashTable()
	excelFileName := "lab8/telephone.xlsx"

	// 打开XLSX文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开文件：", err)
		return nil
	}

	// 遍历所有工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历每一行
		for i, row := range sheet.Rows {
			if i == 0 {
				continue
			}
			ti := telephoneInfo{}
			// 遍历每个单元格
			for col, cell := range row.Cells {
				value := cell.String()
				switch col {
				case 0:
					ti.id, _ = strconv.Atoi(value)
				case 1:
					ti.name = value
				case 2:
					ti.address = value
				case 3:
					ti.phonenum = value
				}
			}
			sc.Put(ti.name, ti)
		}
	}
	return sc
}

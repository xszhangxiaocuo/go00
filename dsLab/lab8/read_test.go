package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "lab8/telephone.xlsx"

	// 打开XLSX文件
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("无法打开文件：", err)
		return
	}

	// 遍历所有工作表
	for _, sheet := range xlFile.Sheets {
		// 遍历每一行
		for _, row := range sheet.Rows {
			// 遍历每个单元格
			for _, cell := range row.Cells {
				value := cell.String()
				fmt.Print(value, " ")
			}
			fmt.Println()
		}
	}
}

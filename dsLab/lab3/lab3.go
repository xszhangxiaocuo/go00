package main

import (
	"dsLab/lab3/Queue"
	"fmt"
	"github.com/xuri/excelize/v2"
)

/*
操作系统的时间片轮转调度算法
自动导出excel表格
2023年4月7日10:36:39
*/

func main() {
	for i := 1; i <= 16; i++ {
		count(uint(i))
	}
}

func count(slices uint) {
	q := createQueue()
	num := q.GetSize()
	result := Queue.NewQueue()
	var time uint = 0
	for !q.IsEmpty() {
		tmp := q.Pop()
		if tmp.T == 0 {
			tmp.Start = time
		}
		time += slices
		tmp.T++
		if tmp.T*slices < tmp.Run {
			q.PushNode(tmp)
		} else {
			tmp.Finish = time
			tmp.Ti = time
			tmp.Wi = float64(time) / float64(tmp.Run)
			result.PushNode(tmp)
		}
	}

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	// 创建一个新的表
	sheet := "Sheet1"
	index, _ := f.NewSheet(sheet)
	// 设置表头
	_ = f.SetCellValue(sheet, "A1", "时间片")
	_ = f.SetCellValue(sheet, "B1", "进程名")
	_ = f.SetCellValue(sheet, "C1", "到达时间")
	_ = f.SetCellValue(sheet, "D1", "运行时间")
	_ = f.SetCellValue(sheet, "E1", "开始时间")
	_ = f.SetCellValue(sheet, "F1", "完成时间")
	_ = f.SetCellValue(sheet, "G1", "周转时间")
	_ = f.SetCellValue(sheet, "H1", "带权周转时间")

	var ave1 uint = 0
	var ave2 float64 = 0
	row := 2

	for !result.IsEmpty() {
		tmp := result.Pop()
		ave1 += tmp.Ti
		ave2 += tmp.Wi

		// 设置单元格值
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), slices)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), tmp.Value)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), tmp.Arrive)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), tmp.Run)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), tmp.Start)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", row), tmp.Finish)
		_ = f.SetCellValue(sheet, fmt.Sprintf("G%d", row), tmp.Ti)
		_ = f.SetCellValue(sheet, fmt.Sprintf("H%d", row), tmp.Wi)

		row++
	}

	// 设置平均周转时间和平均带权周转时间
	_ = f.SetCellValue(sheet, "A"+fmt.Sprint(row), "平均周转时间")
	_ = f.SetCellValue(sheet, "B"+fmt.Sprint(row), float64(ave1)/float64(num))
	_ = f.SetCellValue(sheet, "A"+fmt.Sprint(row+1), "平均带权周转时间")
	_ = f.SetCellValue(sheet, "B"+fmt.Sprint(row+1), ave2/float64(num))

	// 将当前工作表设置为活动工作表
	f.SetActiveSheet(index)
	// 保存到Excel文件中
	err := f.SaveAs(fmt.Sprintf("time_slice_%d.xlsx", slices))
	if err != nil {
		fmt.Println("保存文件失败:", err)
		return
	}

	fmt.Println("文件已保存:", fmt.Sprintf("time_slice_%d.xlsx", slices))
	fmt.Println()
}

func createQueue() *Queue.Queue {
	q := Queue.NewQueue()
	q.Push("A", 0, 20)
	q.Push("B", 0, 10)
	q.Push("C", 0, 15)
	q.Push("D", 0, 5)
	return q
}

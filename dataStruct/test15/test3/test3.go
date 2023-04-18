package main

import "fmt"

/*
运动会设立了n个比赛项目，每个运动员可以参加1~3个项目。请编写程序，安排运动会的比赛日程。要求日程安排既使同一运动员的项目不在同一时间进行，
又使总的竞赛日程尽可能地短。如有多种日程安排，只需找到任意一种即可。
【假设条件】假设运动会有9个项目，项目编号情况为：A={1,2,3,4,5,6,7,8,9}，
9名运动员报名参加的项目分别为：(2,8), (4,5,9), (2,9), (2,1), (2,5,6), (7,5), (7,6), (3,7), (6,3)

2023年4月17日17:11:25
*/

//type Queue struct {
//	proj []int
//	next
//}
//
//func main() {
//	n := 9
//	result := make([]int, n)
//	for i := 0; i < n; i++ {
//
//	}
//}
type Athlete struct {
	id       int
	projects []int
}

func main() {
	athletes := []Athlete{
		{id: 1, projects: []int{2, 8}},
		{id: 2, projects: []int{4, 5, 9}},
		{id: 3, projects: []int{2, 9}},
		{id: 4, projects: []int{2, 1}},
		{id: 5, projects: []int{2, 5, 6}},
		{id: 6, projects: []int{7, 5}},
		{id: 7, projects: []int{7, 6}},
		{id: 8, projects: []int{3, 7}},
		{id: 9, projects: []int{6, 3}},
	}

	athletesInProjects := make(map[int][]int)
	for _, athlete := range athletes {
		for _, project := range athlete.projects {
			athletesInProjects[project] = append(athletesInProjects[project], athlete.id)
		}
	}

	// 项目总数
	projectCount := 9

	// 初始化日程表
	schedule := make([][]int, 0)

	for len(athletesInProjects) > 0 {
		daySchedule := make([]int, 0)
		visitedAthletes := make(map[int]bool)

		for i := 1; i <= projectCount; i++ {
			athletesForProject, ok := athletesInProjects[i]
			if !ok {
				continue
			}

			canAdd := true
			for _, athleteID := range athletesForProject {
				if visitedAthletes[athleteID] {
					canAdd = false
					break
				}
			}

			if canAdd {
				daySchedule = append(daySchedule, i)
				delete(athletesInProjects, i)
				for _, athleteID := range athletesForProject {
					visitedAthletes[athleteID] = true
				}
			}
		}

		schedule = append(schedule, daySchedule)
	}

	// 输出结果
	for i, daySchedule := range schedule {
		fmt.Printf("第%d天：项目%v\n", i+1, daySchedule)
	}
}

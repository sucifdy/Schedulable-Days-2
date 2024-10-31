package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	dateCount := make(map[int]int)

	for _, dates := range villager {
		uniqueDates := make(map[int]struct{})
		for _, date := range dates {
			uniqueDates[date] = struct{}{}
		}

		for date := range uniqueDates {
			dateCount[date]++
		}
	}

	result := []int{}
	for date, count := range dateCount {
		if count == len(villager) {
			result = append(result, date)
		}
	}

	return result
}

func main() {
	villager1 := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(villager1)) // Output: [12, 19]

	villager2 := [][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5}, {2, 3, 4, 10, 11, 12, 15}}
	fmt.Println(SchedulableDays(villager2)) // Output: [2, 3, 4]

	villager3 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {10, 11, 12}, {21, 22, 23, 24}, {25}}
	fmt.Println(SchedulableDays(villager3)) // Output: []

	villager4 := [][]int{}
	fmt.Println(SchedulableDays(villager4)) // Output: []
}

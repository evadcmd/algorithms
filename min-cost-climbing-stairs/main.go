package main

import "fmt"

func min(nums ...int) int {
	if len(nums) < 1 {
		return 0
	}

	m := nums[0]
	for _, i := range nums {
		if m > i {
			m = i
		}
	}

	return m
}

func minCost(costs []int) int {
	l := len(costs)
	switch l {
	case 1:
		return 0
	case 2:
		return min(costs...)
	default:
		return min(
			minCost(costs[:l-1])+costs[l-1],
			minCost(costs[:l-2])+costs[l-2],
		)

	}
}

func main() {
	costs := []int{10, 15, 20}
	fmt.Println(minCost(costs))
	costs2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCost(costs2))
}

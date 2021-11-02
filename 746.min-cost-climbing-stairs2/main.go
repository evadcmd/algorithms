package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCost(cost []int) int {

	l := len(cost)
	if l < 2 {
		return 0
	}

	minCost := make([]int, l+1)
	for i := 2; i < l+1; i++ {
		minCost[i] = min(
			minCost[i-1]+cost[i-1],
			minCost[i-2]+cost[i-2],
		)
	}
	return minCost[l]
}

func main() {
	costs := []int{10, 15, 20}
	fmt.Println(minCost(costs))
	costs2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCost(costs2))
}

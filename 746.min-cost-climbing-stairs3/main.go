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
	for i := 2; i < l; i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[l-1], cost[l-2])
}

func main() {
	costs := []int{10, 15, 20}
	fmt.Println(minCost(costs))
	costs2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCost(costs2))
}

package main

func maxProfit(prices []int) int {
	sum := 0
	for i, diff := 1, 0; i < len(prices); i++ {
		if diff = prices[i] - prices[i-1]; diff > 0 {
			sum += diff
		}
	}
	return sum
}

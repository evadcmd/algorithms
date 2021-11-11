// Use Kadane's Algorithm to separate the array,
// and record the max separately,
// If the sum fall down (max - fee), sell the stock and start a new calculation.
package main

func maxProfit(prices []int, fee int) int {
	sum := 0
	for i, delta, max := 1, 0, 0; i < len(prices); i++ {
		d := prices[i] - prices[i-1]
		delta += d
		if delta > max {
			max = delta
		}
		if delta < 0 || delta-max+fee < 0 || i == len(prices)-1 {
			if profit := max - fee; profit > 0 {
				sum += profit
			}
			delta, max = 0, 0
		}
	}
	return sum
}

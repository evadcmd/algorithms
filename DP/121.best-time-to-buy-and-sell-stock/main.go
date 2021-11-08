// Please read 53.maximum-subarray3 to compare the boundary conditions
package main

func maxProfit(prices []int) int {
	// This question require that if result < 0 return 0,
	// so the default zero value is a better choice.
	max := 0
	for i, sum := 1, 0; i < len(prices); i++ {
		sum += (prices[i] - prices[i-1])
		if sum > max {
			max = sum
			// 53.maximum-subarray allows minus result,
			// so could not allow else if there.
		} else if sum < 0 {
			sum = 0
		}
	}
	return max
}

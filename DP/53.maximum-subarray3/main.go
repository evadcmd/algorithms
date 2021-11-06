package main

import "math"

func maxSubArray(nums []int) int {
	var sum int
	res := math.MinInt
	for _, num := range nums {
		sum += num
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

package main

import "math"

func maxProduct(nums []int) int {
	max := math.MinInt32
	for i := range nums {
		p := 1
		for _, num := range nums[i:] {
			if p = p * num; p > max {
				max = p
			}
		}
	}
	return max
}

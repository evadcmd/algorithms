// brute
package main

import "math"

func maxSubArray(nums []int) int {
	max := math.MinInt32
	for i := range nums {
		m := 0
		for _, num := range nums[i:] {
			m += num
			if m > max {
				max = m
			}
		}
	}
	return max
}

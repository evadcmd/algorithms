package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	res := math.MinInt32
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < i+l; j++ {
			sum += nums[j%l]
			if sum > res {
				res = sum
			}
			if sum < 0 {
				sum = 0
			}
		}
	}
	return res
}

package main

import "math"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	sum := math.MinInt
	for i := 1; i < len(nums); i++ {
		// if the left part is negative, then do not add
		// similar to the part of maximum-subarray3
		// if sum < 0 {
		//     sum = 0
		// }
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		sum = max(sum, nums[i])
	}
	return sum
}

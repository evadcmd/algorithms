package main

import "math"

func maxSubArray(nums []int) int {
	// the default value of min must be zero
	// in first iteration:
	// sum += num <- first value
	// tmp := sum - min (== 0)
	// we get the tmp as first value
	var sum, min int
	res := math.MinInt32
	for _, num := range nums {
		sum += num
		// We intend to put min update after calculate the tmp result
		// because in the case [-a, -b, -c]
		// we could get the correct answer [(-a) + (-b) + (-c)] - [(-a) + (-b)]
		if tmp := sum - min; tmp > res {
			res = tmp
		}
		if sum < min {
			min = sum
		}
	}
	return res
}

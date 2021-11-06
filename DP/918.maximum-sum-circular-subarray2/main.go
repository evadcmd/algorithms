// Please read the 53.maximum-subarray2 to get the concept of this solution.
// At first round, we also record the sum of each index
// in order to compute the min peaks reversely.
// Then we confirm the part of [n, 2n) part of array
// to get the correct max value.
package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	mins := make([]int, l)
	res, sum := math.MinInt32, 0

	for i, min := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		mins[i] = sum
		if r := sum - min; r > res {
			res = r
		}
		if sum < min {
			min = sum
		}
	}

	for i, min := l-2, mins[l-1]; i >= 0; i-- {
		if mins[i] > min {
			mins[i] = min
		} else {
			min = mins[i]
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]
		if r := sum - mins[(i+1)%l]; r > res {
			res = r
		}
	}
	return res
}

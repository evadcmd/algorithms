// We count the stream if the difference between any two consecutive elements is the same,
// and use n * (n-1) / 2 to calculate the result.
// For example:
// * * * * *
// there would be (5 - 3 + 1) = 3 + 2 + 1 = 6 of arithmetic subarrays.
package main

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	res := 0
	count := 1
	prev := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		curr := nums[i] - nums[i-1]
		if curr == prev {
			count++
		} else {
			if count >= 2 {
				res += (count) * (count - 1) >> 1
			}
			count = 1
		}
		prev = curr
	}
	if count >= 2 {
		res += (count) * (count - 1) >> 1
	}
	return res
}

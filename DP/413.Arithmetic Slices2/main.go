// We count the stream if the difference between any two consecutive elements is the same,
// and use n * (n-1) / 2 to calculate the result.
// For example:
// * * * * *
// there would be (5 - 3 + 1) = 3 + 2 + 1 = 6 of arithmetic subarrays.
package main

func numberOfArithmeticSlices(nums []int) int {
	res := 0
	dp := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp++
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}

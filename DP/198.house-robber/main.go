package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	l := len(nums)
	switch l {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	case 3:
		return max(nums[0]+nums[2], nums[1])
	}
	nums[1] = max(nums[0], nums[1])
	nums[2] = max(nums[0]+nums[2], nums[1])
	for i := 3; i < l; i++ {
		nums[i] += max(nums[i-2], nums[i-3])
	}
	return max(nums[l-1], nums[l-2])
}

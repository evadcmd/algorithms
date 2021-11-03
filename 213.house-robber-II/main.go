package main

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

func rob(nums []int) int {
	l := len(nums)
	// special cases: boundary conditions
	switch l {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	case 3:
		return max(nums[0], nums[1], nums[2])
	}
	// choose the 1st house and give up the last house 0 ~ (n - 2)
	cost := make([]int, l)
	cost[0] = nums[0]
	cost[1] = max(nums[0], nums[1])
	cost[2] = max(nums[0]+nums[2], nums[1])
	for i := 3; i < l-1; i++ {
		cost[i] = nums[i] + max(cost[i-2], cost[i-3])
	}
	a, b := cost[l-2], cost[l-3]
	// choose the last house and give up the 1st house 1 ~ (n - 1)
	cost[0] = 0
	cost[1] = nums[1]
	cost[2] = max(nums[2], nums[1])
	for i := 3; i < l; i++ {
		cost[i] = nums[i] + max(cost[i-2], cost[i-3])
	}
	return max(a, b, cost[l-1], cost[l-2])
}

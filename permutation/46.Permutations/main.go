package main

func permute(nums []int) [][]int {
	l := len(nums) - 1
	if l == 0 {
		return [][]int{
			nums,
		}
	}

	if l == 1 {
		return [][]int{
			{nums[0], nums[1]},
			{nums[1], nums[0]},
		}
	}

	res := [][]int{}
	for i := l; i >= 0; i-- {
		nums[l], nums[i] = nums[i], nums[l]
		for _, ns := range permute(nums[:l]) {
			ns = append(ns, nums[l])
			res = append(res, ns)
		}
		nums[i], nums[l] = nums[l], nums[i]
	}
	return res
}

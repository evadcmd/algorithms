package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	/*
		last := len(nums) - 1
		charge := 0
		for i, num := range nums {
			if i == last {
				return true
			}
			charge = max(charge-1, num)
			if charge == 0 {
				break
			}
		}
		return false
	*/
	last := len(nums) - 1
	for i, charge := 0, 0; i <= last && charge >= 0; i, charge = i+1, max(charge, nums[i])-1 {
		if i == last {
			return true
		}
	}
	return false
}

func main() {
	canJump([]int{3, 2, 1, 0, 4})
}

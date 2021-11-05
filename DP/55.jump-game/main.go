package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	l := len(nums)
	for from := l - 2; from >= 0; from-- {
		for step := 1; step <= min(nums[from], l-from-1); step++ {
			if stride := step + nums[from+step]; stride > nums[from] {
				nums[from] = stride
			}
		}
	}
	return nums[0] >= (l - 1)
}

func main() {
	canJump([]int{3, 2, 1, 0, 4})
}

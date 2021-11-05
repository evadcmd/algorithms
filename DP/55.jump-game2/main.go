package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func canJump(nums []int) bool {
	l := len(nums)
	possible := make([]bool, l)
	possible[l-1] = true
	for from := l - 2; from >= 0; from-- {
		// the index which is bigger has more possiblility to reach
		for step := min(nums[from], l-from-1); step > 0; step-- {
			if possible[from+step] {
				possible[from] = true
				break
			}
			possible[from] = false
		}
	}
	return possible[0]
}

func main() {
	canJump([]int{3, 2, 1, 0, 4})
}

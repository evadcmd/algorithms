package main

import "math"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func jump(nums []int) int {
	l := len(nums)

	jmp := make([]int, l)
	for i := 1; i < l; i++ {
		jmp[i] = math.MaxInt32
	}

	for i, v := range nums {
		for j := i + 1; j <= min(i+v, l-1); j++ {
			if c := jmp[i] + 1; c < jmp[j] {
				jmp[j] = c
			}
		}
	}

	return jmp[l-1]
}

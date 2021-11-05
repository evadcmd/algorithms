package main

import "math"

func jump(nums []int) int {
	l := len(nums)
	jmp := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		// if it could reach the last index directly, it costs 1 step.
		if nums[i] >= l-1-i {
			jmp[i] = 1
			continue
		}
		// find a min step from [i+1, i+nums[i]+1)
		jmp[i] = math.MaxInt32
		/*
			for _, v := range jmp[i+1 : i+nums[i]+1] {
				if s := v + 1; s < jmp[i] {
					jmp[i] = s
				}
			}
		*/
		for s := i + nums[i]; s > i; s-- {
			if j := jmp[s] + 1; j < jmp[i] {
				jmp[i] = j
			}
		}
	}
	return jmp[0]
}

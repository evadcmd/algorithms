// a greedy approch of the jump game
package main

func jump(nums []int) int {
	// becareful with the boundary condition
	// for example: nums := [0] <= need zero jumps to reach the last index
	jmp := 0
	for i, curr, next := 0, 0, 0; i < len(nums); i++ {
		// default value: i == 0, curr == 0
		// Would not be executed on first iteration.

		// On 2nd iteration, the 「curr」 must be updated first
		// because we want to put first round step of nums[0].
		// If we changed the order of these two If statements.
		// the first round step would be set to nums[1]
		// which is not correct.
		if i > curr {
			curr = next
			jmp++
		}
		if limit := i + nums[i]; limit > next {
			next = limit
		}
	}
	return jmp
}

package main

func climbStairs(n int) int {
	a, b := 1, 1
	for n--; n > 0; n-- {
		b += a
		a = b - a
	}
	return b
}

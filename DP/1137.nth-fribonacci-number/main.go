package main

func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}

	a, b, c := 0, 1, 1
	for n -= 2; n > 0; n-- {
		c = a + b + c
		b = c - a - b
		a = c - b - a
	}
	return c
}

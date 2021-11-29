package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for n--; n > 0; n-- {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(fib(5))
}

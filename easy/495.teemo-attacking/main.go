package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	for i := 1; i < len(timeSeries); i++ {
		total += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return total + duration
}

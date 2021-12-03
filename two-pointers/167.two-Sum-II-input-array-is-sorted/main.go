package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := target + 1
	for sum != target {
		sum = numbers[l] + numbers[r]
		if sum > target {
			r--
		} else if sum < target {
			l++
		}
	}
	return []int{l + 1, r + 1}
}

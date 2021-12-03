package main

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	l := len(arr)
	diff := arr[l-1] - arr[0]
	res := [][]int{}
	for i := 1; i < len(arr); i++ {
		dif := arr[i] - arr[i-1]
		if dif < diff {
			diff = dif
			res = [][]int{}
		}
		if dif == diff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}

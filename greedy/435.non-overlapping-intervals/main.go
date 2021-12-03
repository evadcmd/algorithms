package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count, bound := 0, intervals[0][1]
	for _, interval := range intervals[1:] {
		if interval[0] < bound {
			count++
		} else {
			bound = interval[1]
		}
	}
	return count
}

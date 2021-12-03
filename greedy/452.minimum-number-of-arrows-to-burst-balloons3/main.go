// sort by start
package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count, bound := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > bound {
			bound = points[i][1]
			count++
		}
	}
	return count
}

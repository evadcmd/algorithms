// be care of the boundary:
// start boundary: default space := 1 to assure the empty spaces to be calcaulated correctly.
// end boundary: n := space / 2 rather then (space - 1) / 2
package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	space := 1
	for _, flower := range flowerbed {
		switch {
		case n <= 0:
			break
		case flower == 0:
			space++
		case flower == 1:
			n -= ((space - 1) >> 1)
			space = 0
		}
	}
	n -= (space >> 1)
	return n <= 0
}

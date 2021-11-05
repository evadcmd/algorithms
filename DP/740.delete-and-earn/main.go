package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Try to interate only if the key exists but is has some problems to be solved.
/*
func deleteAndEarn2(nums []int) int {
    l := len(nums)
    earn := make(map[int]int, l)
    for _, num := range nums {
        earn[num] += num
    }
    keys := []int{}
    for k := range earn {
        keys = append(keys, k)
    }
    if len(keys) == 1 {
        return earn[keys[0]]
    }
    sort.Ints(keys)
    for _, k := range keys {
        earn[k] += max(earn[k-2], earn[k-3]) // not correct because the value may not be filled in there (k-2, k-3)
    }
    to1 := keys[len(keys)-1]
    to2 := keys[len(keys)-2]
    return max(earn[to1], earn[to2])
}
*/

func deleteAndEarn(nums []int) int {
	earn := make(map[int]int, len(nums))
	from, to := nums[0], nums[0]
	for _, num := range nums {
		earn[num] += num
		if num < from {
			from = num
		}
		if num > to {
			to = num
		}
	}
	for i := from; i <= to; i++ {
		earn[i] += max(earn[i-2], earn[i-3])
	}
	return max(earn[to-1], earn[to])
}

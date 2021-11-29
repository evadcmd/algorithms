package main

import "fmt"

func nextPermutation(nums []int) {
	l := len(nums) - 1
	i := l
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			for j := i; j < len(nums); j++ {
				if nums[j] <= nums[i-1] {
					nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
					break
				}
				if j == l {
					nums[i-1], nums[j] = nums[j], nums[i-1]
				}
			}
			break
		}
	}
	fmt.Println(nums)
	sub := nums[i:]
	for k := 0; k < len(sub)>>1; k++ {
		sub[k], sub[len(sub)-1-k] = sub[len(sub)-1-k], sub[k]
	}

	/*
	   for k := 0; (i + k) < (l + i) >> 1; k++ {
	       nums[i+k], nums[l-k] = nums[l-k], nums[i+k]
	   }
	*/

}

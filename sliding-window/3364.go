package main

import "math"

func minimumSumSubarray(nums []int, l int, r int) int {
	minR := math.MaxInt

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			jarak := j - i + 1

			if jarak >= l && jarak <= r && sum > 0 {
				minR = min(minR, sum)
			}
		}
	}

	if minR == math.MaxInt {
		return -1
	}

	return minR
}

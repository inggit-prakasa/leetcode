package main

import "math"

// Bruteforce
func maxRotateFunction_brute_force(nums []int) int {
	res := math.MinInt
	arrResp := make([]int, len(nums))

	for i, _ := range arrResp {
		for j, num := range nums {
			temp := (j - i + len(nums)) % len(nums)
			arrResp[i] += temp * num
		}
	}

	for _, arr := range arrResp {
		if arr > res {
			res = arr
		}
	}

	return res
}

func maxRotateFunction(nums []int) int {
	sum := 0
	n := len(nums)

	for _, num := range nums {
		sum += num
	}

	max := 0
	for i, num := range nums {
		max += i * num
	}

	res := max
	if n > 1 {
		for i := 1; i < n; i++ {
			currMax := max + sum - (n * nums[n-i])
			if currMax > res {
				res = currMax
			}

			max = currMax
		}
	}

	return res
}

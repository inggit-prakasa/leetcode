package main

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/description/?envType=daily-question&envId=2025-11-12
func minOperations(nums []int) int {
	n := len(nums)

	if slices.Contains(nums, 1) {
		return n - 1
	}

	for start := 0; start < n; start++ {
		currentGCD := nums[start]
		for end := start + 1; end < n; end++ {
			currentGCD = gcd(currentGCD, nums[end])
			if currentGCD == 1 {
				return n - 1 + (end - start)
			}
		}
	}

	return -1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func trialDivision(a int) bool {
	if a < 2 {
		return false
	}

	if a == 2 {
		return true
	}

	if a%2 == 0 {
		return false
	}

	sqrtInt := int(math.Sqrt(float64(a)))
	for i := 3; i <= sqrtInt; i += 2 {
		if a%i == 0 {
			return false
		}
	}

	return true
}

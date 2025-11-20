package main

// https://leetcode.com/problems/keep-multiplying-found-values-by-two/?envType=daily-question&envId=2025-11-19
func findFinalValue(nums []int, original int) int {
	resp := original

	for i := 0; i < len(nums); i++ {
		if nums[i] == resp {
			resp *= 2
			i = -1
		}
	}

	return resp
}

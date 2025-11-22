package main

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/?envType=daily-question&envId=2025-11-22
func minimumOperations(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	resp := 0
	for _, num := range nums {
		if num%3 != 0 {
			resp++
		}	
	}

	return resp
}

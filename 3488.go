package main

import "sort"

// Bruteforce approach
func ClosestEqualElement_BruteForce(queries []int, nums []int) []int {
	res := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		temp := -1

		num := nums[queries[i]]
		min := len(nums)

		for j := 0; j < len(nums); j++ {
			if num == nums[j] && queries[i] != j {
				minA := abs(queries[i] - j)
				minB := len(nums) - minA

				if minA < min {
					min = minA
				}

				if minB < min {
					min = minB
				}

			}

			if min != len(nums) {
				temp = min
			}
		}

		res[i] = temp
	}

	return res
}

//
func solveQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	temp := make(map[int][]int)

	for idx, num := range nums {
		temp[num] = append(temp[num], idx)
	}

	for idx, q := range queries {
		if len(temp[nums[q]]) <= 1 {
			res[idx] = -1
			continue
		}

		x := sort.SearchInts(temp[nums[q]], q)

		lTemp := len(temp[nums[q]])
		leftX := abs(temp[nums[q]][x] - temp[nums[q]][(x-1+lTemp)%lTemp])
		left := min(leftX, len(nums)-leftX)

		rightX := abs(temp[nums[q]][x] - temp[nums[q]][(x+1)%lTemp])
		right := min(rightX, len(nums)-rightX)

		res[idx] = min(left, right)
	}

	return res
}

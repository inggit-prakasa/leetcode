package main

func specialTriplets(nums []int) int {
	resp := 0
	left := map[int]int{}
	right := map[int]int{}
	modMax := 1_000_0007
	for _, num := range nums {
		right[num] = right[num] + 1
	}

	for j := 0; j < len(nums); j++ {
		right[nums[j]] = right[nums[j]] - 1
		target := nums[j] * 2

		countLeft := 0
		countRight := 0

		if left[target] > 0 {
			countLeft = left[target]
		}

		if right[target] > 0 {
			countRight = right[target]
		}

		resp += (countLeft * countRight) % modMax

		left[nums[j]] = left[nums[j]] + 1
	}

	return resp
}

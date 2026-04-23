package main

// Brute force solution
func distance_brute_force(nums []int) []int64 {
	res := make([]int64, len(nums))
	for i := 0; i < len(nums); i++ {

		temp := 0
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i] == nums[j] {
				temp += abs(i - j)
			}
		}

		res[i] = int64(temp)
	}
	return res
}

func distance(nums []int) []int64 {
	res := make([]int64, len(nums))

	totalSum := make(map[int]int64)
	totalCount := make(map[int]int)

	for i, num := range nums {
		totalSum[num] += int64(i)
		totalCount[num]++
	}

	prefixSum := make(map[int]int64)
	prefixCount := make(map[int]int)

	for i, num := range nums {

		leftDist := int64(prefixCount[num]*i) - prefixSum[num]

		rightDist := totalSum[num] - int64(i*totalCount[num])

		totalSum[num] -= int64(i)
		totalCount[num]--

		res[i] = leftDist + rightDist

		prefixSum[num] += int64(i)
		prefixCount[num]++
	}

	return res
}

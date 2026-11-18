package main

func majorityElement(nums []int) int {
	temp := make(map[int]int, 0)
	for _, v := range nums {
		temp[v]++
	}

	maxTemp := 0
	maxKey := nums[0]
	for key, v := range temp {
		if v > maxTemp {
			maxTemp = v
			maxKey = key
		}
	}

	return maxKey
}

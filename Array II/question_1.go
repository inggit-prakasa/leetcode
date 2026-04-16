package main

func findErrorNums(nums []int) []int {
	mapNums := make([]int, len(nums)+1)
	for _, num := range nums {
		mapNums[num]++
	}

	dup, miss := -1, -1
	for i := 1; i <= len(nums); i++ {
		if mapNums[i] == 2 {
			dup = i
		} else if mapNums[i] == 0 {
			miss = i
		}
	}

	return []int{dup, miss}
}

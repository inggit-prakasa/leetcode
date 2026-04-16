package main

func repeatedNTimes(nums []int) int {

	mapNums := map[int]int{}
	for _, num := range nums {
		mapNums[num]++
	}

	max := 0
	key := 0
	for i, num := range mapNums {
		if num > max {
			max = num
			key = i
		}
	}

	return key
}

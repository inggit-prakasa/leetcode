package main

func minimumPairRemoval(nums []int) int {
	ans := 0
	sum := 0
	tempI := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if i == 0 {
			tempI++
			continue
		}

		if nums[i] >= nums[i-1] {
			tempI++
		}
	}

	if tempI == len(nums) {
		return 0
	}

	tempNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			tempNums[i] = nums[i]
			continue
		}

		tempNums[i] = sum - nums[i-1]
		sum -= nums[i-1]
	}

	temp := 0
	for i := 0; i < len(nums)-1; i++ {
		if tempNums[i] > tempNums[i+1] {
			temp = i
			break
		}
	}

	ans = len(nums) - temp - 1
	return ans
}

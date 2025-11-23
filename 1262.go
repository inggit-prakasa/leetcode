package main

// https://leetcode.com/problems/greatest-sum-divisible-by-three/description/?envType=daily-question&envId=2025-11-23
func maxSumDivThree(nums []int) int {

	temp := map[int][]int{}
	sum := 0

	for _, num := range nums {
		sum += num
		mod := num % 3
		temp[mod] = append(temp[mod], num)
	}

	for k := range temp {
		temp[k] = sortArrayList(temp[k])
	}

	mod := sum % 3
	switch mod {
	case 1:
		option1 := 0
		option2 := 0
		if len(temp[1]) >= 1 {
			option1 = sum - temp[1][0]
		}

		if len(temp[2]) >= 2 {
			option2 = sum - temp[2][0] - temp[2][1]
		}

		sum = max(option1, option2)
	case 2:
		option1 := 0
		option2 := 0

		if len(temp[2]) >= 1 {
			option1 = sum - temp[2][0]
		}

		if len(temp[1]) >= 2 {
			option2 = sum - temp[1][0] - temp[1][1]
		}

		sum = max(option1, option2)
	}

	return sum
}

// DP Approach
// func maxSumDivThree(nums []int) int {
// 	// dp[i] = max sum dengan remainder i
// 	dp := [3]int{0, math.MinInt32, math.MinInt32}

// 	for _, num := range nums {
// 		temp := dp
// 		for i := 0; i < 3; i++ {
// 			newRem := (i + num) % 3
// 			dp[newRem] = max(dp[newRem], temp[i]+num)
// 		}
// 	}

// 	return dp[0]
// }

func sortArrayList(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

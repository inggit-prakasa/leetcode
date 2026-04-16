package main

func maxDotProduct(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := 0; i <= len(nums1); i++ {
		dp[i][0] = -1001
	}

	for j := 0; j <= len(nums2); j++ {
		dp[0][j] = -1001
	}

	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			a := nums1[i-1] * nums2[j-1]
			b := nums1[i-1]*nums2[j-1] + dp[i-1][j-1]
			c := dp[i-1][j]
			d := dp[i][j-1]
			dp[i][j] = max4(a, b, c, d)
		}
	}

	return dp[len(nums1)][len(nums2)]
}

func max4(a, b, c, d int) int {
	return max(max(a, b), max(c, d))
}

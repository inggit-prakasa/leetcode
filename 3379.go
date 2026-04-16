package main

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		x := nums[i]

		y := (i + x) % n

		if y < 0 {
			ans[i] = nums[y+n]
			continue
		}

		ans[i] = nums[y]
	}

	return ans
}

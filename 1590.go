package main

func minSubarray(nums []int, p int) int {
	minSum := 1_000_000_007

	sum := 0
	for _, num := range nums {
		sum += num
	}

	targerMod := sum % p
	if targerMod == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			subSum := 0
			for k := i; k <= j; k++ {
				subSum += nums[k]
			}

			if (sum-subSum)%p == 0 {
				if j-i+1 < minSum {
					minSum = j - i + 1
				}
			}
		}
	}




	return minSum
}
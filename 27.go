package main

func removeElement(nums []int, val int) int {
	ln := 0
	for _, v := range nums {
		if v != val {
			nums[ln] = v
			ln++
		}
	}

	return ln
}

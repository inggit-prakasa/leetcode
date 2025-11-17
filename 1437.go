package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	x := 0
	resp := false
	for _, num := range nums {
		if num == 1 {
			if x < k {
				resp = false
			} else {
				resp = true
			}
			x = 0
			continue
		}

		x++
	}

	fmt.Println("x dan k", x, k)

	if x >= k - 1 {
		return true
	}

	return resp
}

package main

func prefixesDivBy5(nums []int) []bool {
	resp := make([]bool, len(nums))

	num := 0
	for a, b := range nums {
		num = (num*2 + b) % 5
		if num == 0 {
			resp[a] = true
		} else {
			resp[a] = false
		}
	}

	return resp
}

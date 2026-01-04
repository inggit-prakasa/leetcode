package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			return digits
		} else if digits[i] == 9 {
			digits[i] = 0
		}
	}

	resp := make([]int, len(digits)+1)
	resp[0] = 1

	return resp
}

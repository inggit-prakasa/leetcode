package main

import "fmt"

// https://leetcode.com/problems/1-bit-and-2-bit-characters
func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}

	i := 0
	n := len(bits)
	for i < n-1 {
		if bits[i] == 1 {
			i += 2
		} else if bits[i] == 0 {
			i += 1
		}
	}

	fmt.Println(i)
	if i == n-1 {
		return true
	}

	return false
}

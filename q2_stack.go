package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	ans := 0
	nums := []int{}
	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		if err != nil {
			sum := 0
			numA := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			numB := nums[len(nums)-1]
			nums = nums[:len(nums)-1]

			fmt.Println(numA, numB, tokens[i])
			switch tokens[i] {
			case "+":
				sum = numB + numA
			case "*":
				sum = numB * numA
			case "-":
				sum = numB - numA
			case "/":
				sum = numB / numA
			}

			nums = append(nums, sum)
		} else {
			nums = append(nums, num)
		}
	}

	ans = nums[0]
	return ans
}

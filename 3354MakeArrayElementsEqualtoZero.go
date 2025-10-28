package main

import "fmt"

func main() {
	req := []int{16, 13, 10, 0, 0, 0, 10, 6, 7, 8, 7}
	fmt.Println(countValidSelections(req))
}

func countValidSelectionsOLD(nums []int) int {
	validSelection := 0
	currs := []int{}
	for i, num := range nums {
		if num == 0 {
			currs = append(currs, i)
		}
	}

	for _, curr := range currs {
		numsLeft := make([]int, len(nums))
		copy(numsLeft, nums)

		mins := "-"
		idx := curr

		for {
			if numsLeft[idx] == 0 {
				switch mins {
				case "-":
					idx--
				case "+":
					idx++
				}

				if idx < 0 || idx >= len(numsLeft) {
					break
				}
			} else if numsLeft[idx] > 0 {
				numsLeft[idx] -= 1
				switch mins {
				case "-":
					mins = "+"
				case "+":
					mins = "-"
				}

				switch mins {
				case "+":
					idx++
				case "-":
					idx--
				}
			}
		}

		check := len(numsLeft)
		for _, v := range numsLeft {
			if v == 0 {
				check--
			}

			if check == 0 {
				validSelection++
				break
			}
		}
	}

	for _, curr := range currs {
		numsLeft := make([]int, len(nums))
		copy(numsLeft, nums)

		mins := "+"
		idx := curr
		for {
			if numsLeft[idx] == 0 {
				switch mins {
				case "-":
					idx--
				case "+":
					idx++
				}

				if idx < 0 || idx >= len(numsLeft) {
					break
				}

			} else if numsLeft[idx] > 0 {

				numsLeft[idx] -= 1
				switch mins {
				case "-":
					mins = "+"
				case "+":
					mins = "-"
				}

				switch mins {
				case "+":
					idx++
				case "-":
					idx--
				}
			}
		}

		check := len(numsLeft)
		for _, v := range numsLeft {
			if v == 0 {
				check--
			}

			if check == 0 {
				validSelection++
				break
			}
		}
	}

	return validSelection
}

func countValidSelections(nums []int) int {
	validSelection := 0
	currs := []int{}
	for i, num := range nums {
		if num == 0 {
			currs = append(currs, i)
		}
	}

	for _, curr := range currs {
		left := 0
		for i := curr; i >= 0; i-- {
			left += nums[i]
		}

		right := 0
		for i := curr + 1; i < len(nums); i++ {
			right += nums[i]
		}

		if left == right {
			validSelection += 2
		} else if left+1 == right || right+1 == left {
			validSelection += 1
		}
	}

	return validSelection
}

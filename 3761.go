package main

// Bruteforce approach
func minMirrorPairDistance_Bruteforce(nums []int) int {
	res := -1
	temp := len(nums)

	for i, num := range nums {
		newNum := reverse(num)
		for j := i + 1; j < len(nums); j++ {
			if newNum == nums[j] {
				temp = min(temp, j-i)
				res = temp
				break
			}
		}
	}
	return res
}

// With hash map
func minMirrorPairDistance(nums []int) int {
	res := -1
	temp := len(nums)
	m := make(map[int]int)

	for i, num := range nums {
		_, ok := m[num]
		if ok {
			temp = min(temp, i-m[num])
			res = temp
		}

		m[reverse(num)] = i
	}

	return res
}

func reverse(num int) int {
	res := 0
	for num != 0 {
		z := num % 10
		num = (num - z) / 10
		res = res*10 + z
	}

	return res
}

// 123 % 10 = 3
// 123 - 3 = 120
// 120 / 10 = 12

// 0 * 10 + 3 = 3
// 3 * 10 + 2 = 32
// 32 * 10 + 1 = 321

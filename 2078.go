package main

// Brute force approach: O(n^2)
func maxDistance_brute_force(colors []int) int {
	res := 0
	temp := make(map[int][]int)

	for i := 1; i < len(colors); i++ {
		temp[colors[i]] = append(temp[colors[i]], i)
	}

	for key, nums := range temp {
		for i := 0; i < len(colors); i++ {
			if colors[i] != key {
				for _, num := range nums {
					res = max(res, abs(num-i))
				}
			}
		}
	}

	return res
}

// Matematical approach: O(n)
func maxDistance(colors []int) int {
	res := 0

	for i := len(colors) - 1; i >= 0; i-- {
		if colors[0] != colors[i] {
			res = max(res, abs(i))
			break
		}
	}

	for i := 0; i < len(colors); i++ {
		if colors[len(colors)-1] != colors[i] {
			res = max(res, abs(len(colors)-i-1))
			break
		}
	}

	return res
}

package main

// Brute force approach: O(n^2)
func maxDistance(colors []int) int {
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

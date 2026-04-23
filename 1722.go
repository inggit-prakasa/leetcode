package main

// Brute force approach: Generate all possible permutations of the source array based on the allowed swaps and calculate the Hamming distance for each permutation. This approach is not efficient and may not be feasible for large input sizes.
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {

	res := 0
	swp := make(map[int][]int)
	for _, nn := range allowedSwaps {
		swp[nn[0]] = append(swp[nn[0]], nn[1])
		swp[nn[1]] = append(swp[nn[1]], nn[0])
	}

	for idxNum, num := range source {

		targetIdx := 0
		for idx, i := range target {
			if num == i {
				targetIdx = idx
				break
			}
		}

		if targetIdx == idxNum {
			continue
		}

		visited := map[int]bool{}
		if !dfs_hamming(idxNum, targetIdx, swp, visited) {
			res++
		}
	}

	return res
}

func dfs_hamming(idx, targetIdx int, mapSwap map[int][]int, visited map[int]bool) bool {
	visited[idx] = true

	if idx == targetIdx {
		return true
	}

	for _, num := range mapSwap[idx] {
		vis, _ := visited[num]
		if vis {
			continue
		}

		if dfs_hamming(num, targetIdx, mapSwap, visited) {
			return true
		}
	}

	return false
}

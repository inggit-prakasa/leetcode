package main

// https://leetcode.com/problems/set-intersection-size-at-least-two/?envType=daily-question&envId=2025-11-20
func intersectionSizeTwo(intervals [][]int) int {
	sortIntervals := sortArray(intervals)

	resp := []int{}
	for i := 0; i < len(sortIntervals); i++ {
		interval := sortIntervals[i]

		if len(resp) == 0 {
			resp = append(resp, interval[1]-1)
			resp = append(resp, interval[1])
			continue
		}

		start := interval[0]
		end := interval[1]

		count := 0
		for _, num := range resp {
			if num >= start && num <= end {
				count++
			}
		}

		// Tambah angka sesuai kebutuhan
		if count >= 2 {
			continue
		} else if count == 1 {
			resp = append(resp, end)
		} else {
			resp = append(resp, end-1)
			resp = append(resp, end)
		}
	}

	return len(resp)
}

func sortArray(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][1] > intervals[j+1][1] ||
				(intervals[j][1] == intervals[j+1][1] && intervals[j][0] < intervals[j+1][0]) {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	return intervals
}

func contains(slice []int, v int) bool {
	for _, x := range slice {
		if x == v {
			return true
		}
	}
	return false
}

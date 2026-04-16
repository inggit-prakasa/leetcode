package main

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	for i := 0; i < len(points)-1; i++ {
		a := abs(points[i+1][0] - points[i][0])
		b := abs(points[i+1][1] - points[i][1])

		ans += max(a, b)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

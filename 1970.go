package main

func latestDayToCross(row int, col int, cells [][]int) int {
	left, right := 1, row*col
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if canCrossBottom(row, col, mid, cells) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

type Point struct {
	r, c int
}

func canCrossBottom(row, col, day int, cells [][]int) bool {

	grid := make([][]int, row)
	for i := range grid {
		grid[i] = make([]int, col)
	}

	for i := 0; i < day; i++ {
		r, c := cells[i][0]-1, cells[i][1]-1
		grid[r][c] = 1
	}

	// queue
	queue := []Point{}
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}

	// add multi bfs, with row 0
	for c := 0; c < col; c++ {
		if grid[0][c] == 0 {
			queue = append(queue, Point{0, c})
			visited[0][c] = true
		}
	}

	// direction
	direction := [4]Point{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:] // Dequeue

		// Kalau nyampe baris paling bawah, WIN!
		if curr.r == row-1 {
			return true
		}

		// Cek tetangga
		for _, d := range direction {
			nr, nc := curr.r+d.r, curr.c+d.c

			// Validasi: Dalam peta, Belum dikunjungi, dan DARATAN
			if nr >= 0 && nr < row && nc >= 0 && nc < col && !visited[nr][nc] && grid[nr][nc] == 0 {
				visited[nr][nc] = true
				queue = append(queue, Point{nr, nc})
			}
		}
	}

	return false
}

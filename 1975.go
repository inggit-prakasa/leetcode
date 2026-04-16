package main

func maxMatrixSum(matrix [][]int) int64 {
	ans := int64(0)
	sum := 0
	temp := 0
	num := 1_000_001
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			tempNum := matrix[i][j]
			if matrix[i][j] < 0 {
				sum = sum - matrix[i][j]
				tempNum = -matrix[i][j]
				temp++
			} else {
				sum += matrix[i][j]
			}

			if tempNum < num {
				num = tempNum
			}
		}
	}

	if temp%2 != 0 {
		ans = int64(sum - 2*num)
	} else {
		ans = int64(sum)
	}

	return ans
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

	func maxProduct(root *TreeNode) int {

		arrSum := make([]int64, 0)

		var dfs func(root *TreeNode) int64
		dfs = func(root *TreeNode) int64 {
			var sum int64 = 0

			if root == nil {
				return sum
			}

			left := dfs(root.Left)
			right := dfs(root.Right)
			val := int64(root.Val)

			sum = val + left + right

			arrSum = append(arrSum, sum)
			return sum
		}

		var totalSum int64 = dfs(root)
		var maxSum int64 = 0
		for _, sum := range arrSum {
			m := sum * (totalSum - sum)

			if m > maxSum {
				maxSum = m
			}
		}

		return int(maxSum % 1000000007)
	}

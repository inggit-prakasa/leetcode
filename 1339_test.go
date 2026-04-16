package main

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: 110,
		},
		{
			name: "balanced small",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 4},
			},
			expected: 20,
		},
		{
			name: "all ones",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 1},
			},
			expected: 2,
		},
		{
			name: "different ",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 10},
				Right: &TreeNode{Val: 2},
			},
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProduct(tt.root)
			if result != tt.expected {
				t.Errorf("maxProduct(%v) = %v; want %v", tt.root, result, tt.expected)
			}
		})
	}
}

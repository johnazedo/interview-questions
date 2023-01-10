package medium

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Difficult: Medium
Link: https://leetcode.com/problems/binary-tree-level-order-traversal/
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var matrix [][]int
	queue := []*TreeNode{root}
	level := 0

	// Time: O(N)
	// Space: O(N)
	for len(queue) > 0 {
		matrix = append(matrix, []int{})
		for _, root := range queue {
			matrix[level] = append(matrix[level], root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
			queue = queue[1:]
		}
		level++
	}
	return matrix
}

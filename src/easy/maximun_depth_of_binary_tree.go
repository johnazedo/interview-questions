package easy

import (
	. "github.com/johnazedo/leetcode/src"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/maximum-depth-of-binary-tree
*/
func maxDepth(root *TreeNode) int {
	// TODO: Do this iteratively
	// Time: O(N)
	// Space: O(H)
	return recursiveMaxDepth(root, 1)
}

func recursiveMaxDepth(root *TreeNode, level int) int {
	if root == nil {
		return level - 1
	}

	left := recursiveMaxDepth(root.Left, level+1)
	right := recursiveMaxDepth(root.Right, level+1)

	if left >= right {
		return left
	} else {
		return right
	}
}

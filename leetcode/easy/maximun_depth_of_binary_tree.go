package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Number: 104
Difficult: Easy
Link: https://leetcode.com/problems/maximum-depth-of-binary-tree
Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
Status: Resolved
*/
func maxDepth(root *TreeNode) int {
	// TODO: Do this iteratively
	return recursiveMaxDepth(root, 1)
}

// Recursive approach
func recursiveMaxDepth(root *TreeNode, level int) int {
	// Time: O(N)
	// Space: O(H) -> H is the height of the tree
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

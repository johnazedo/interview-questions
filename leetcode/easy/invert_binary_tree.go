package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Easy
LinK: https://leetcode.com/problems/invert-binary-tree/
*/
func invertTree(root *TreeNode) *TreeNode {
	// TODO: Try this iteratively
	// Time: O(N)
	// Space: O(H)
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

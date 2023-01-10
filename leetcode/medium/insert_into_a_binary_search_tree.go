package medium

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Medium
Link: https://leetcode.com/problems/insert-into-a-binary-search-tree/
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// TODO: Can create this recursively?
	// Time: O(Log(N))
	// Space: O(1)
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for curr != nil {
		if curr.Val < val {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &TreeNode{Val: val}
				break
			}
		}
	}
	return root
}

package medium

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Difficult: Medium
Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// TODO: Can you do this iteratively?
	// Time: O(N)
	// Space: O(H)
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if root.Val == p.Val {
		return p
	} else if root.Val == q.Val {
		return q
	}

	if left == nil && right == nil {
		return nil
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	} else {
		return root
	}
}

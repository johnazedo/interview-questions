package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Number: 100
Difficult: Easy
Link: https://leetcode.com/problems/same-tree/
Tags: Binary Tree, Depth-First Search
Status: Resolved
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// TODO: Do this iteratively
	// Time: O(N)
	// Space: O(H)
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

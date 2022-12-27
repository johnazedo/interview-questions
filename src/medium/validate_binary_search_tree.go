package medium

import (
	. "github.com/johnazedo/leetcode/src"
	"math"
)

/*
Difficult: Medium
Link: https://leetcode.com/problems/validate-binary-search-tree/
*/
func isValidBST(root *TreeNode) bool {
	min := math.MinInt
	max := math.MaxInt
	// Time: O(N)
	// Space: O(N)
	return helper(root, min, max)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val < min || root.Val > max {
		return false
	}

	return helper(root.Left, min, root.Val-1) &&
		helper(root.Right, root.Val+1, max)
}

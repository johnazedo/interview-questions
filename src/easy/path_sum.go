package easy

import . "github.com/johnazedo/leetcode/src"

/*
Difficult: Easy
Link: https://leetcode.com/problems/path-sum/
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	// TODO: Try this iteratively
	// Time: O(N)
	// Space: O(H)
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)

}

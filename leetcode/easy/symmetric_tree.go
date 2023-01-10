package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Easy
Link: https://leetcode.com/problems/symmetric-tree/
*/
func isSymmetric(root *TreeNode) bool {
	// TODO: Do this iteratively
	// Time: O(N)
	// Space: O(H)
	return recursiveIsSymmetric(root.Left, root.Right)
}

func recursiveIsSymmetric(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	out := recursiveIsSymmetric(left.Left, right.Right)
	in := recursiveIsSymmetric(left.Right, right.Left)

	return out && in && left.Val == right.Val
}

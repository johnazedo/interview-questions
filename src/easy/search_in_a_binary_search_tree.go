package easy

import . "github.com/johnazedo/leetcode/src"

/*
Difficult: Easy
Link: https://leetcode.com/problems/search-in-a-binary-search-tree/
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	// Time: O(Log(N))
	// Space: O(1)
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}
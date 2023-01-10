package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Easy
Link: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
*/
func findTarget(root *TreeNode, k int) bool {
	// Time: O(N)
	// Space: O(N + H)
	mapping := make(map[int]int) // Space O(N)
	stack := []*TreeNode{root}   // Space O(H)

	for len(stack) > 0 {
		n := len(stack) - 1
		root = stack[n]
		stack = stack[:n]
		if root != nil {
			if _, ok := mapping[k-root.Val]; ok {
				return true
			} else {
				mapping[root.Val] = 0
				stack = append(stack, root.Left, root.Right)
			}
		}
	}

	return false
}

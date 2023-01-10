package easy

import (
	. "github.com/johnazedo/leetcode/leetcode"
)

/*
Number: 145
Difficult: Easy
Link: https://leetcode.com/problems/binary-tree-postorder-traversal/
Tags: Stack, Tree, Depth-First Search
Status: Resolved
*/
func postorderTraversal(root *TreeNode) []int {
	// N is the number of nodes and H is the height of the tree
	// Time: O(N)
	// Space: O(H)
	var result []int
	recursivePost(root, &result)
	return result
}

func recursivePost(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	recursivePost(root.Left, result)
	recursivePost(root.Right, result)
	*result = append(*result, root.Val)
}

// Has no official solution by leetcode.
// This is not working
// Error: out of memory (infinite loop)
// Solution: Use a right child stack
func iterativePost(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	var lastNode *TreeNode

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			n := len(stack) - 1
			peekNode := stack[n]
			if peekNode.Right != nil && lastNode != peekNode.Right {
				root = peekNode.Right
			} else {
				stack = stack[:n]
				result = append(result, peekNode.Val)
				lastNode = root
			}
		}
	}
	return result
}

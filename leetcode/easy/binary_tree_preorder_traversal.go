package easy

import . "github.com/johnazedo/leetcode/leetcode"

/*
Number: 144
Difficult: Easy
Link: https://leetcode.com/problems/binary-tree-preorder-traversal/
Tags: Stack, Tree, Depth-First Search
Status: Reviewed
*/
func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var result []int
	// N is the number of nodes and H is the height of the tree
	// Time: O(N)
	// Space: O(H)
	for len(stack) > 0 {
		n := len(stack) - 1
		root = stack[n]
		stack = stack[:n]
		if root != nil {
			result = append(result, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return result
}

func recursivePre(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	recursivePre(root.Left, result)
	recursivePre(root.Right, result)
}

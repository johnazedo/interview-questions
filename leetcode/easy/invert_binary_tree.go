package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Number 226
Difficult: Easy
LinK: https://leetcode.com/problems/invert-binary-tree/
Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
Status: Reviewed
*/
func invertTree(root *TreeNode) *TreeNode {
	// Recursive solution
	// Time: O(N)
	// Space: O(H)
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTreeIT(root *TreeNode) *TreeNode {
	// Iterative solution
	// Time: O(N)
	// Space: O(H)

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return root
}

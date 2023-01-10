package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/n-ary-tree-preorder-traversal
*/

func preorder(root *Node) []int {
	result := make([]int, 0)
	stack := make([]*Node, 0)

	stack = append(stack, root)
	for len(stack) > 0 {
		n := len(stack) - 1
		root = stack[n]
		if root != nil {
			result = append(result, root.Val)
			stack = stack[:n]
			for i := len(root.Children) - 1; i >= 0; i-- {
				stack = append(stack, root.Children[i])
			}
		}
	}
	return result
}

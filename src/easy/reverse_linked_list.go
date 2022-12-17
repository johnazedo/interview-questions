package easy

import "github.com/johnazedo/leetcode/src"

/*
Difficult: Easy
Link: https://leetcode.com/problems/reverse-linked-list
*/
func reverseList(head *src.ListNode) *src.ListNode {
	// Time: O(N)
	// Space: O(N)

	stack := []int{}
	temp := head

	for temp != nil {
		stack = append(stack, temp.Val)
		temp = temp.Next
	}

	temp = head

	for i := len(stack) - 1; i >= 0; i-- {
		temp.Val = stack[i]
		temp = temp.Next
	}

	return head
}

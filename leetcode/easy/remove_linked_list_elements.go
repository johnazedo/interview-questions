package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/remove-linked-list-elements/
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// Time: O(N)
	// Space: O(1)
	var prevNode *ListNode
	temp := head
	for temp != nil {
		if val == temp.Val {
			if prevNode != nil {
				prevNode.Next = temp.Next
				temp = temp.Next
			} else {
				head = head.Next
				temp = head
			}
		} else {
			prevNode = temp
			temp = temp.Next
		}
	}
	return head
}

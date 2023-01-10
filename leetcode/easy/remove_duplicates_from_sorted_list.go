package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Difficult: Easy
Link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head

	// Time: O(N)
	// Space: O(1)
	for temp != nil && temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	return head
}

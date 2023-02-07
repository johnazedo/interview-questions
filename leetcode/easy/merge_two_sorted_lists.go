package easy

import . "github.com/johnazedo/interview-questions/leetcode"

/*
Number: 21
Difficult: Easy
Link: https://leetcode.com/problems/merge-two-sorted-lists
Tags: Linked List, Recursion
Status: Reviewed
*/

// Iteratively approach
func mergeTwoListsIt(list1 *ListNode, list2 *ListNode) *ListNode {
	// Time: O(M+N)
	// Space: O(1)
	head := &ListNode{}
	var v1 int
	var v2 int
	temp := head
	for list1 != nil || list2 != nil {
		if list1 != nil {
			v1 = list1.Val
		} else {
			v1 = 999
		}
		if list2 != nil {
			v2 = list2.Val
		} else {
			v2 = 999
		}

		if v1 <= v2 {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next

	}
	return head.Next
}

// Recursive approach
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Time O(M+N)
	// Space O(M+N)
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

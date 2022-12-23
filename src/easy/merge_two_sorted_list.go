package easy

import "github.com/johnazedo/leetcode/src"

/*
Difficult: Easy
Link: https://leetcode.com/problems/merge-two-sorted-lists
*/
func mergeTwoLists(list1 *src.ListNode, list2 *src.ListNode) *src.ListNode {
	head := &src.ListNode{}
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
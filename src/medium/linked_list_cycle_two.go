package medium

import (
	"github.com/johnazedo/leetcode/src"
)

/*
Difficult: Medium
Link: https://leetcode.com/problems/linked-list-cycle-ii
*/
func detectCycle(head *src.ListNode) *src.ListNode {
	mapping := make(map[*src.ListNode]int)
	for {
		if head == nil {
			return nil
		}
		if _, ok := mapping[head]; ok {
			break
		}
		mapping[head] = 0
		head = head.Next
	}

	return head
}

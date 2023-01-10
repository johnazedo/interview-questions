package medium

import (
	"github.com/johnazedo/interview-questions/leetcode"
)

/*
Difficult: Medium
Link: https://leetcode.com/problems/linked-list-cycle-ii
*/
func detectCycle(head *leetcode.ListNode) *leetcode.ListNode {
	mapping := make(map[*leetcode.ListNode]int)
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

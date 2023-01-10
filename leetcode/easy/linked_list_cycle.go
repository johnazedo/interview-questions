package easy

import (
	. "github.com/johnazedo/leetcode/leetcode"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/linked-list-cycle/
*/
func hasCycle(head *ListNode) bool {

	// TODO: Do this with O(1) space complexity
	// Time: O(N)
	// Space: O(N)
	mapping := make(map[*ListNode]int)
	if head == nil {
		return false
	}

	for head != nil {
		if _, ok := mapping[head]; ok {
			return true
		}
		mapping[head] = 0
		head = head.Next
	}

	return false
}

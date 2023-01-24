package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
)

/*
Number 141
Difficult: Easy
Link: https://leetcode.com/problems/linked-list-cycle/
Tags: Hash Table, Linked List, Two Pointers
Status: Resolved
*/

// HashTable approach
func hasCycle(head *ListNode) bool {
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

// Two pointers approach
func hasCycleBetterSpaceComplexity(head *ListNode) bool {
	// Time: O(N) -> Because the worst case is when the slowest one does a full loop
	// Space: O(1)

	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

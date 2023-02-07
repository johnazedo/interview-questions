package easy

import (
	. "github.com/johnazedo/interview-questions/leetcode"
	"math"
)

/*
Number: 876
Difficult: Easy
Link: https://leetcode.com/problems/middle-of-the-linked-list
Tags: Linked List, Two Pointers
Status: Reviewed
*/

func middleNode(head *ListNode) *ListNode {
	// Time: O(N)
	// Space: O(1)
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	middle := head

	for head.Next != nil {
		count++
		head = head.Next
	}

	r := math.Ceil(float64(count / 2))

	for i := 0; i < int(r); i++ {
		middle = middle.Next
	}

	return middle
}

// Best approach: Slow and Fast pointer
func middleNodeBest(head *ListNode) *ListNode {
	// Time: O(N)
	// Space: O(1)
	slow := head
	for head != nil && head.Next != nil {
		slow = slow.Next
		head = head.Next.Next
	}

	return slow
}

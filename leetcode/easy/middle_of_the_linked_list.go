package easy

import (
	"github.com/johnazedo/interview-questions/leetcode"
	"math"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/middle-of-the-linked-list
*/
func middleNode(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
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

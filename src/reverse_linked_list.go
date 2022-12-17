package src

func reverseList(head *ListNode) *ListNode {
	stack := []int{}
	temp := head

	for temp != nil {
		stack = append(stack, temp.Val)
		temp = temp.Next
	}

	temp = head

	for i := len(stack) - 1; i >= 0; i-- {
		temp.Val = stack[i]
		temp = temp.Next
	}

	return head
}

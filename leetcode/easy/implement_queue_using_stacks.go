package easy

/*
Number: 232
Difficult: Easy
Link: https://leetcode.com/problems/implement-queue-using-stacks/
Tags: Stack, Design, Queue
Status: Reviewed
*/

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.stack = append(q.stack, x)
}

func (q *MyQueue) Pop() int {
	value := q.stack[0]
	q.stack = q.stack[1:]
	return value
}

func (q *MyQueue) Peek() int {
	return q.stack[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.stack) == 0
}

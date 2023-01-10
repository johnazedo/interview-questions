package easy

/*
Number: 70
Difficult: Easy
Link: https://leetcode.com/problems/climbing-stairs/
Tags: Math, Dynamic Programming, Memorization
Status: Reviewed
*/
func climbStairs(n int) int {
	// Time: O(N)
	// Space: O(1)
	prev := 1
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/climbing-stairs/
*/
func climbStairs(n int) int {
	// TODO: Do this with dynamic programming
	// Time: O(N)
	// Space: O(1)
	prev := 1
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

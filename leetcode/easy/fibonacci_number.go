package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/fibonacci-number/
*/
func fib(n int) int {
	// Time: O(N)
	// Space: O(1)
	if n < 2 {
		return n
	}
	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

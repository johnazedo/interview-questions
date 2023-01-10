package easy

/*
Number 509
Difficult: Easy
Link: https://leetcode.com/problems/fibonacci-number/
Tags: Math, Dynamic Programming
Status: Reviewed
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

func fibDynamic(n int) int {
	// Time: O(1) OBS: If dp array was given
	// Space: O(N)
	if n == 0 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

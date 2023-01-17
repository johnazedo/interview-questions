package easy

/*
Number: 392
Difficult: Easy
Link: https://leetcode.com/problems/is-subsequence
Tags: Two Pointers, String, Dynamic Programming
Status: Resolved
*/
func isSubsequence(s string, t string) bool {
	// TODO: Do this with dynamic programming
	// TODO: Do this with divider to conquer approach
	sPtr := 0
	tPtr := 0

	// Two pointers solution
	// Time: O(N)
	// Space: O(1)
	// OBS: N is len(t)
	for sPtr < len(s) {
		if tPtr >= len(t) {
			return false
		}
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
		tPtr++
	}

	return true
}

package src

/*
Difficult: Easy
Link: https://leetcode.com/problems/is-subsequence
*/
func isSubsequence(s string, t string) bool {
	sPtr := 0
	tPtr := 0

	// O(N) no qual N é o tamanho de t
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

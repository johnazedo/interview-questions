package src

import "testing"

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

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		Params   []string
		Expected bool
	}{
		{Params: []string{"abc", "ahbgdc"}, Expected: true},
		{Params: []string{"axc", "ahbgdc"}, Expected: false},
	}

	for _, test := range tests {
		result := isSubsequence(test.Params[0], test.Params[1])
		if result != test.Expected {
			t.Fatalf("Expected value: %v - Actual: %v", test.Expected, result)
		}
	}
}

package easy

/*
Number: 205
Difficult: Easy
Link: https://leetcode.com/problems/isomorphic-strings
Tags: Hash Table, String
Status: Reviewed
*/
func isIsomorphic(s string, t string) bool {
	// Time: O(N)
	// Space: O(N+M)
	stot := make(map[byte]byte)
	ttos := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sValue, sInT := stot[s[i]]
		tValue, tInS := ttos[t[i]]

		if !sInT && !tInS {
			stot[s[i]] = t[i]
			ttos[t[i]] = s[i]
		} else if sValue != t[i] || tValue != s[i] {
			return false
		}
	}

	return true
}

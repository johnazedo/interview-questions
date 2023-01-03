package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/backspace-string-compare/
*/
func backspaceCompare(s string, t string) bool {
	// Time: O(N + M)
	// Space: O(N + M)
	return getNewString(s) == getNewString(t)
}

func getNewString(s string) string {
	var runes []rune

	for _, r := range s {
		if r == '#' {
			n := len(runes)
			if n != 0 {
				runes = runes[:n-1]
			}
		} else {
			runes = append(runes, r)
		}
	}

	return string(runes)
}

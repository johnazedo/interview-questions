package easy

import (
	"strings"
)

/*
Difficult: Easy
Link: https://leetcode.com/problems/length-of-last-word/
*/
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	lastSpace := -1

	// Time: O(N)
	// Space: O(1)
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			lastSpace = i
			break
		}
	}

	return len(s[lastSpace+1:])
}

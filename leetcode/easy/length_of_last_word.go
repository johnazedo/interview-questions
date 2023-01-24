package easy

import (
	"strings"
)

/*
Number: 58
Difficult: Easy
Link: https://leetcode.com/problems/length-of-last-word/
Tags: String
Status: Reviewed
*/
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	lastSpace := -1

	// Time: O(N)
	// Space: O(N) -> Because Trim makes a copy of the original string
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			lastSpace = i
			break
		}
	}

	return len(s[lastSpace+1:])
}

// One loop approach
func lengthOfLastWordBetterSpaceComplexity(s string) int {
	// Time: O(N)
	// Space: O(1)

	pointer := len(s) - 1
	length := 0

	for pointer >= 0 {

		if string(s[pointer]) != " " {
			length++
		} else if length > 0 {
			break
		}

		pointer--
	}

	return length
}

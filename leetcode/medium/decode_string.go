package medium

import (
	"strconv"
	"strings"
)

/*
Number: 394
Difficult: Medium
Link: https://leetcode.com/problems/decode-string/
Tags: String, Stack, Recursion
Status: Unsolved
*/
func decodeString(s string) string {
	/*
		Help: ASCII Table

		[ => 91
		] => 93
		0 - 9 => 48 - 57
		a - z => 97 - 122
	*/

	result := ""
	var cStack = ""
	var nStack = ""

	for i, c := range s {

		if c >= 48 && c <= 57 {
			nStack = nStack + string(c)
		}

		if c >= 97 && c <= 122 {
			cStack = cStack + string(c)
		}

		if c == 91 {
			result += decodeString(s[i+1:])
			break
		}
	}

	value, err := strconv.Atoi(nStack)
	if err != nil || value == 0 {
		return cStack
	}
	result = strings.Repeat(result, value)

	return cStack + result
}

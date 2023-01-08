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
    for i, c := range s {
        if string(c) == " " {
            lastSpace = i
        }
    }

    return len(s[lastSpace+1:])
}

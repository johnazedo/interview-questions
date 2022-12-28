package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/valid-parentheses/
*/
func isValid(s string) bool {
	var stack []rune
	mapping := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	// Time: O(N)
	// Space: O(N)
	for _, c := range s {
		n := len(stack) - 1
		if value, ok := mapping[c]; ok && n >= 0 {
			if value != stack[n] {
				return false
			}
			stack = stack[:n]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

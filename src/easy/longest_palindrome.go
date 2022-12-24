package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/longest-palindrome/
*/
func longestPalindrome(s string) int {
	mapping := make(map[string]int)
	sum := 0
	for _, c := range s {
		key := string(c)
		if _, ok := mapping[key]; ok {
			sum += 2
			delete(mapping, key)
		} else {
			mapping[key] = 1
		}
	}

	if len(mapping) > 0 {
		sum++
	}
	return sum
}

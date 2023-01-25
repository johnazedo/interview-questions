package easy

/*
Number: 409
Difficult: Easy
Link: https://leetcode.com/problems/longest-palindrome/
Tags: Hash Table, String, Greedy
Status: Reviewed
*/
func longestPalindrome(s string) int {
	// Time: O(N)
	// Space: O(1) -> Because there are only 26 letters in English alphabet
	mapping := make(map[rune]int)
	sum := 0
	for _, c := range s {
		// When found a duplicate letter, add 2 to counter and remove this from map.
		if _, ok := mapping[c]; ok {
			sum += 2
			delete(mapping, c)
		} else {
			mapping[c] = 1
		}
	}

	// If the size of map is grater than 0 it means that
	// you can choose one letter to put in middle of palindrome,
	// so add 1 to counter.
	if len(mapping) > 0 {
		sum++
	}
	return sum
}

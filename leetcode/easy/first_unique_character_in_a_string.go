package easy

/*
Number: 387
Difficult: Easy
Link: https://leetcode.com/problems/first-unique-character-in-a-string/
Tags: Hash Table, String, Queue
Status: Reviewed
*/
func firstUniqChar(s string) int {
	// Time: O(N)
	// Space: O(1) -> Because there are only 26 letters in english alphabet
	mapping := make(map[rune]int)
	for _, c := range s {
		mapping[c] = mapping[c] + 1
	}

	for i, c := range s {
		if mapping[c] == 1 {
			return i
		}
	}

	return -1
}

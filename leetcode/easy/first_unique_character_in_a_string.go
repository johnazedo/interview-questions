package easy

/*
Number: 387
Difficult: Easy
Link: https://leetcode.com/problems/first-unique-character-in-a-string/
Tags: Hash Table, String, Queue
Status: Resolved
*/
func firstUniqChar(s string) int {
	// TODO: Review this question
	// Time: O(N)
	// Space: O(N)
	mapping := make(map[int32]int)
	for i, c := range s {
		if _, ok := mapping[c]; ok {
			mapping[c] = -1
			continue
		}
		mapping[c] = i
	}

	min := 100000
	for _, v := range mapping {
		if v < min && v != -1 {
			min = v
		}
	}

	if min == 100000 {
		return -1
	}

	return min
}

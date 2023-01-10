package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/valid-anagram/
*/
func isAnagram(s string, t string) bool {
	mapping := make(map[int32]int)
	// N = len(s) and M = len(t)
	// Time: O(N+M)
	// Space: O(N)
	for _, c := range s {
		mapping[c] += 1
	}

	for _, c := range t {
		if _, ok := mapping[c]; !ok {
			return false
		}
		mapping[c] -= 1
		if mapping[c] == 0 {
			delete(mapping, c)
		}
	}

	return len(mapping) == 0
}

package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/ransom-note/
*/
func canConstruct(ransomNote string, magazine string) bool {
	mapping := make(map[int32]int)

	// N = len(ransomNote) and M = len(magazine)
	// Time: O(N+M)
	// Space: O(N)
	for _, c := range ransomNote {
		mapping[c] += 1
	}

	for _, c := range magazine {
		if _, ok := mapping[c]; ok {
			mapping[c] -= 1
			if mapping[c] == 0 {
				delete(mapping, c)
			}
		}
	}

	return len(mapping) == 0
}

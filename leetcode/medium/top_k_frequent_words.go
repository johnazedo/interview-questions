package medium

import "sort"

/*
Number: 692
Difficult: Medium
Link: https://leetcode.com/problems/top-k-frequent-words/
Tags: Hash Table, String, Trie, Sorting, Heap(Priority Queue), Bucket Sort, Counting
Status: Resolved
*/
func topKFrequent(words []string, k int) []string {
	// TODO:  Could you solve it in O(n log(k)) time and O(n) extra space?
	// Time: O(K*N)
	// Space: O(N)
	mapping := make(map[string]int)
	var result []string

	for _, s := range words {
		mapping[s] = mapping[s] + 1
		result = append(result, s)
	}

	sort.Strings(result) // O(N*LogN)

	// O(K*N)
	for i := 0; i < k; i++ {
		max := 0
		word := ""
		for _, s := range result {
			if mapping[s] > max {
				max = mapping[s]
				word = s
			}
		}
		result = append(result, word)
		delete(mapping, word)
	}

	return result[len(result)-k:]
}

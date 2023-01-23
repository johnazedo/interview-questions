package easy

import (
	"sort"
)

/*
Number: 1046
Difficult: Easy
Link: https://leetcode.com/problems/last-stone-weight/
Tags: Array, Heap (Priority Queue)
Status: Resolved
*/

// Sorting array and get the largest approach
func lastStoneWeight(stones []int) int {
	// Time: O(Nˆ2 * Log(N))
	// Space: O(1)
	for len(stones) > 1 {
		sort.Ints(stones) // O(N*LogN)
		n := len(stones) - 1
		y := stones[n]
		x := stones[n-1]
		stones = stones[:n-1]
		if y > x {
			stones = append(stones, y-x)
		} else {
			stones = append(stones, x-y)
		}
	}

	return stones[0]
}

// Heap-Based simulation approach
func lastStoneWeightHeapBased(stones []int) int {
	// TODO: Implement a priority queue
	return 0
}

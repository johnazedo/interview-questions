package easy

import "sort"

/*
Difficult: Easy
Link: https://leetcode.com/problems/last-stone-weight/
*/
func lastStoneWeight(stones []int) int {
	// TODO: Check time complexity
	// Time: O(Nˆ2)
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

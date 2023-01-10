package easy

/*
Number 217
Difficult: Easy
Link: https://leetcode.com/problems/contains-duplicate/
Tags: Array, Hash Table, Sorting
Status: Reviewed
*/
func containsDuplicate(nums []int) bool {
	mapping := make(map[int]int)

	// Time: O(N)
	// Space: O(N)
	for _, n := range nums {
		if _, ok := mapping[n]; ok {
			return true
		}
		mapping[n] = 0
	}

	return false
}

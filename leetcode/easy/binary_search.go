package easy

/*
Number: 704
Difficult: Easy
Link: https://leetcode.com/problems/binary-search/
Tags: Array, Binary Search
Status: Reviewed
*/
func search(nums []int, target int) int {
	// Time: O(LogN)
	// Space: O(1)
	start := 0
	end := len(nums) - 1
	var pivot int
	for start <= end {
		pivot = start + (end-start)/2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] < target {
			start = pivot + 1
		}
		if nums[pivot] > target {
			end = pivot - 1
		}
	}
	return -1
}

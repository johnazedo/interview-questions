package easy

// Example
func isBadVersion(version int) bool { return false }

/*
Number: 278
Difficult: Easy
Link: https://leetcode.com/problems/first-bad-version/
Tags: Binary Search, Interactive
Status: Reviewed
*/
func firstBadVersion(n int) int {
	// Time: O(LogN)
	// Space: O(1)
	start := 1
	end := n
	var pivot int = start + (end-start)/2
	for start <= end {
		if isBadVersion(pivot) {
			end = pivot - 1
		} else {
			start = pivot + 1
		}
		pivot = start + (end-start)/2
	}
	return pivot
}

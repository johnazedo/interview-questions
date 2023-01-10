package medium

/*
Difficult: Medium
Link: https://leetcode.com/problems/maximum-subarray/
*/
func maxSubArray(nums []int) int {
	total := 0
	max := -9999999

	// Time: O(N^2)
	// Space: O(1)
	for x, _ := range nums {
		total = 0
		for y, ny := range nums {
			if y < x {
				continue
			}
			total += ny
			if total > max {
				max = total
			}
		}
	}
	return max
}

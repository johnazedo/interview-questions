package medium

/*
Number: 53
Difficult: Medium
Link: https://leetcode.com/problems/maximum-subarray/
Tags: Array, Divider and Conquer, Dynamic Programming
Status: Resolved
*/

// Brute force approach (not accepted)
func maxSubArray(nums []int) int {
	total := 0
	max := -9999999
	// TODO: Do this with divide to conquer approach

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

// Dynamic programming approach
func maxSubArrayDP(nums []int) int {
	// Time: O(N)
	// Space: O(1)
	maxSub := nums[0]
	currentSub := 0

	for _, n := range nums {
		currentSub += n
		if currentSub > maxSub {
			maxSub = currentSub
		}

		if currentSub < 0 {
			currentSub = 0
		}
	}
	return maxSub
}

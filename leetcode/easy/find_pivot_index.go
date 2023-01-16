package easy

/*
Number: 724
Difficult: Easy
Link: https://leetcode.com/problems/find-pivot-index
Tags: Array, Prefix Sum
Status: Reviewed
*/
func pivotIndex(nums []int) int {
	// Time: O(N)
	// Space: O(1)
	mount := 0
	leftSum := 0
	for _, n := range nums {
		mount += n
	}
	for i := 0; i < len(nums); i++ {
		if leftSum == mount-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

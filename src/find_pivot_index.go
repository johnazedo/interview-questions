package src

/*
Difficult: Easy
Link: https://leetcode.com/problems/find-pivot-index
*/
func pivotIndex(nums []int) int {
	mount := 0 // O(1)
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			mount += nums[i-1] // O(N)
		}
		sum := 0

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j] // O(N^2)
		}
		if sum == mount { // O(N)
			return i // O(1)
		}
	}

	return -1
}

package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/find-pivot-index
*/
func pivotIndex(nums []int) int {
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

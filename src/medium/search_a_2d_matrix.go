package medium

/*
Difficult: Medium
Link: https://leetcode.com/problems/search-a-2d-matrix/
*/
func searchMatrix(matrix [][]int, target int) bool {
	// TODO: Do this with O(Log N+M) complexity where N = len(matrix) and M = len(matrix[0)
	// Time: O(N) -> N = len(matrix)
	// Space: O(1)
	for i := 0; i < len(matrix); i++ {
		if i == len(matrix)-1 {
			return binarySearch(matrix[i], target)
		}
		if matrix[i][0] <= target && matrix[i+1][0] > target {
			return binarySearch(matrix[i], target)
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	var pivot int
	for start <= end {
		pivot = start + (end-start)/2
		if nums[pivot] == target {
			return true
		} else if nums[pivot] < target {
			start = pivot + 1
		} else {
			end = pivot - 1
		}
	}
	return false
}

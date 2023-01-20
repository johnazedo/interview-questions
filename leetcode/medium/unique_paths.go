package medium

/*
Number: 62
Difficult: Medium
Link: https://leetcode.com/problems/unique-paths/
Tags: Math, Dynamic Programming, Combinatorics
Status: Resolved
*/
func uniquePaths(m int, n int) int {
	// TODO: Do this with O(M) space complexity
	// Time: O(MxN)
	// Space: O(MxN)
	var matrix [][]int

	for i := 0; i < m; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				matrix[i] = append(matrix[i], 1)
			} else {
				matrix[i] = append(matrix[i], matrix[i-1][j]+matrix[i][j-1])
			}
		}
	}
	return matrix[m-1][n-1]
}

package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/pascals-triangle/
*/
func generate(numsRows int) [][]int {
	// Time: O(N^2)
	// Space: O(N^2)
	m := make([][]int, numsRows)
	for i := 0; i < numsRows; i++ {
		m[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				m[i][j] = 1
			} else {
				m[i][j] = m[i-1][j-1] + m[i-1][j]
			}
		}
	}
	return m
}

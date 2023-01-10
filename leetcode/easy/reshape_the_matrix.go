package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/reshape-the-matrix/
*/

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	matrix := make([][]int, r)
	oi := 0
	oj := 0

	// TODO: Do this faster than O(R*C)
	for i := 0; i < r; i++ {
		matrix[i] = make([]int, c)
		for j := 0; j < c; j++ {
			if oj == len(mat[0]) {
				oi++
				oj = 0
			}
			matrix[i][j] = mat[oi][oj]
			oj++
		}
	}

	return matrix
}

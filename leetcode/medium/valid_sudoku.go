package medium

/*
Number: 36
Difficult: Medium
Link: https://leetcode.com/problems/valid-sudoku/description/
Tags: Array, Hash Table, Matrix
Status: Resolved
*/
func isValidSudoku(board [][]byte) bool {
	// TODO: Do this with O(Nˆ2) time complexity
	// TODO: Do this with Bitmasking
	// Time: O(Nˆ3)
	// Space: O(Nˆ2)
	mapping := make(map[byte][][2]int)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			number := board[i][j]
			if number != 46 {
				if value, ok := mapping[number]; ok {
					if !checkIndexes(value, i, j) {
						return false
					}
				}
				// Here, we save the positions that number appears
				mapping[number] = append(mapping[number], [2]int{i, j})
			}
		}
	}
	return true
}

func checkIndexes(vector [][2]int, x, y int) bool {
	// For each position saved of number, we check if the new position
	// are in the same column, row or box.
	for i := 0; i < len(vector); i++ {
		if vector[i][0] == x || vector[i][1] == y || (x/3 == vector[i][0]/3 && y/3 == vector[i][1]/3) {
			return false
		}
	}
	return true
}

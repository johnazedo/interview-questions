package medium

/*
Number: 200
Difficult: Medium
Link: https://leetcode.com/problems/number-of-islands/description/
Tags: Array, Depth-First Search, Breadth-First Search, Union Find
Status: Resolved
*/

func numIslands(grid [][]byte) int {
	// TODO: Do this with Breadth-First Search
	// TODO: Do this with Union Find (Disjoint Set)
	// Time: O(Nˆ2)
	// Space: O(1)

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 49 {
				count++
				changeIsland(&grid, i, j)
			}
		}
	}
	return count
}

func changeIsland(grid *[][]byte, r, c int) {
	if (*grid)[r][c] == 49 {
		(*grid)[r][c] = 48

		if r >= 1 {
			changeIsland(grid, r-1, c)
		}
		if c >= 1 {
			changeIsland(grid, r, c-1)
		}
		if r+1 < len(*grid) {
			changeIsland(grid, r+1, c)
		}
		if c+1 < len((*grid)[0]) {
			changeIsland(grid, r, c+1)
		}
	}
}

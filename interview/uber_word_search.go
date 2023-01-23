package interview

/*
INTERVIEW DATE: December 05, 2022
COMPANY: Uber
STATUS: Resolved

DESCRIPTION:
Given a two-dimensional array of letters, find a given word written in any of the 8 directions.
I.e.

EXAMPLE
Input: UBER
Input:
A  U  I  K  F  W  N
W  Q  B  O  L  X  P
T  L  A  E  R  E  S
Y  Z  X  E  R  L  W
Output: true
*/

func WordSearch(s string, m [][]uint8) bool {
	// Time: O(N^2)
	// Space: O(M) where M = len(s)
	count := 0
	for x := 0; x < len(m); x++ {
		for y := 0; y < len(m[0]); y++ {
			if s[count] == m[x][y] {
				if search(x-1, y-1, -1, -1, m, s, count+1) || // up left
					search(x-1, y, -1, 0, m, s, count+1) || // up
					search(x-1, y+1, -1, 1, m, s, count+1) || // up right
					search(x, y+1, 0, 1, m, s, count+1) || // right
					search(x+1, y+1, 1, 1, m, s, count+1) || // down right
					search(x+1, y, 1, 0, m, s, count+1) || // down
					search(x+1, y-1, 1, -1, m, s, count+1) || // down left
					search(x, y-1, 0, -1, m, s, count+1) { // left
					return true
				}
			}
		}
	}
	return false
}

func search(x, y, i, j int, m [][]uint8, s string, count int) bool {
	if x >= 0 && x < len(m) && y >= 0 && y < len(m[0]) {
		if m[x][y] == s[count] {
			if count == len(s)-1 {
				return true
			} else {
				return search(x+i, y+i, i, j, m, s, count+1)
			}
		}
	}
	return false
}

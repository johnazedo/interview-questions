package easy

/*
Number: 733
Difficult: Easy
Link: https://leetcode.com/problems/flood-fill/
Tags: Array, Depth-First Search, Breadth-First Search
Status: Resolved
*/
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// TODO: Do this iteratively
	// Time: O(N)
	// Space: O(N)
	prev := image[sr][sc]
	if prev != color {
		floodFillRec(&image, sr, sc, color, prev)
	}
	return image
}

func floodFillRec(image *[][]int, sr, sc, color, prev int) {
	if (*image)[sr][sc] == prev {
		(*image)[sr][sc] = color

		if sr >= 1 {
			floodFillRec(image, sr-1, sc, color, prev)
		}
		if sc >= 1 {
			floodFillRec(image, sr, sc-1, color, prev)
		}
		if sr+1 < len(*image) {
			floodFillRec(image, sr+1, sc, color, prev)
		}
		if sc+1 < len((*image)[0]) {
			floodFillRec(image, sr, sc+1, color, prev)
		}
	}
}

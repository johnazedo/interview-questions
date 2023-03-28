package abc294

import (
	"fmt"
)

/*
Difficult: B
Score: 200
Link: https://atcoder.jp/contests/abc294/tasks/abc294_b
Contest: abc294
*/

func asciiArt() {
	// Time: O(N^2)
	// Space: O(1)

	var rows, cols, temp int
	fmt.Scanf("%d %d", &rows, &cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&temp)
			if temp == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf(string(rune(temp + 64)))
			}
		}
		fmt.Print("\n")
	}
}

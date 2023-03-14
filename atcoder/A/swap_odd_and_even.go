package A

import "fmt"

/*
Difficult: A
Score: 100
Link: https://atcoder.jp/contests/abc293/tasks/abc293_a
Contest: abc293
*/

func main() {
	// Time: O(N)
	// Space: O(N)

	var input, output string
	var save rune
	fmt.Scanf("%s", &input)

	for i, c := range input {
		if (i+1)%2 != 0 {
			save = c
		} else {
			output += string(c)
			output += string(save)
		}
	}

	fmt.Println(output)
}

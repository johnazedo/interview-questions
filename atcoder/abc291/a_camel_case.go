package abc291

import "fmt"

/*
Difficult: A
Score: 100
Link: https://atcoder.jp/contests/abc291/tasks/abc291_a?lang=en
Contest: abc291
*/

func main() {
	// Time: O(N)
	// Space: O(1)

	var text string
	fmt.Scanf("%s", &text)

	for i, c := range text {
		if c < 97 {
			fmt.Println(i + 1)
		}
	}
}

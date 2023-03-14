package A

import "fmt"

/*
Difficult: A
Score: 100
Link: https://atcoder.jp/contests/abc270/tasks/abc270_a?lang=en
Contest: abc270
*/

func main() {
	// Time: O(1)
	// Space: O(1)
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a | b)
}

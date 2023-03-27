package abc295

import (
	"fmt"
)

/*
Difficult: A
Score: 100
Link: https://atcoder.jp/contests/abc295/tasks/abc295_a
Contest: abc295
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var s = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s[i])
	}

	var mapping = map[string]int{
		"and":  1,
		"not":  2,
		"that": 3,
		"the":  4,
		"you":  5,
	}

	for _, r := range s {
		if _, ok := mapping[r]; ok {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}

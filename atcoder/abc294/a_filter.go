package abc294

import (
	"fmt"
)

/*
Difficult: A
Score: 100
Link: https://atcoder.jp/contests/abc294/tasks/abc294_a
Contest: abc294
*/

func main() {
	// Time: O(N)
	// Space: O(N)

	var amount int
	fmt.Scan("%d", &amount)

	arr := make([]int, amount)
	for i := 0; i < amount; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for _, v := range arr {
		if v%2 == 0 {
			fmt.Printf("%d ", v)
		}
	}
}

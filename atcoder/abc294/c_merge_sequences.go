package abc294

import "fmt"

/*
Difficult: C
Score: 300
Link: https://atcoder.jp/contests/abc294/tasks/abc294_c
Contest: abc294
*/

func main() {

	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	vectorA := make([]int, a)
	vectorB := make([]int, b)

	for i := 0; i < a; i++ {
		fmt.Scan(&vectorA[i])
	}

	for i := 0; i < b; i++ {
		fmt.Scan(&vectorB[i])
	}

}

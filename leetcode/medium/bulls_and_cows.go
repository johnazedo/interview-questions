package medium

import "fmt"

/*
Number: 299
Difficult: Medium
Link: https://leetcode.com/problems/bulls-and-cows/
Tags: Hash Table, String, Counting
Status: Unsolved
*/
func getHint(secret string, guess string) string {
	mapping := make(map[rune][]int)

	for i, c := range secret {
		if _, ok := mapping[c]; ok {
			mapping[c] = append(mapping[c], i)
		} else {
			mapping[c] = []int{i}
		}
	}

	bulls := 0
	cows := 0

	for i, c := range guess {
		if value, ok := mapping[c]; ok {
			cows++
			for _, v := range value {
				if i == v {
					bulls++
					cows--
				}
			}
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

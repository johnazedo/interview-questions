package easy

import "math"

/*
Difficult: Easy
Link: https://leetcode.com/problems/min-cost-climbing-stairs/
*/
func minCostClimbingStairs(cost []int) int {
	// TODO: Review this question
	// Time: O(N)
	// Space: O(1)
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] = cost[i] + int(math.Min(float64(cost[i+1]), float64(cost[i+2])))
	}

	return int(math.Min(float64(cost[0]), float64(cost[1])))
}

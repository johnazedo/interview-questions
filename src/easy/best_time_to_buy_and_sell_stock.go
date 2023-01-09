package easy

/*
Number: 121
Difficult: Easy
Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
Tags: Array, Dynamic Programming
Status: Reviewed
*/
func maxProfit(prices []int) int {
	left := 0
	right := 1
	max := 0
	// Time: O(N)
	// Space: O(1)
	for left < len(prices) && right < len(prices) {
		value := prices[right] - prices[left]
		if value <= 0 {
			left++
		} else {
			if value > max {
				max = value
			}
			right++
		}
		if left == right {
			right++
		}
	}

	return max
}

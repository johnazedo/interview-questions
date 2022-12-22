package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/
func maxProfit(prices []int) int {

	left := 0
	right := 1
	max := 0

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

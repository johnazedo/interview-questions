package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for i, n := range nums {
		if index, ok := mapping[target-n]; ok {
			return []int{i, index}
		}
		mapping[n] = i
	}

	return nil
}

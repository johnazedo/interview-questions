package easy

import "sort"

/*
Difficult: Easy
Link: https://leetcode.com/problems/intersection-of-two-arrays-ii/
*/
func intersect(nums1 []int, nums2 []int) []int {
	// Time: O(n*log(n))
	// Space: O(1)

	// TODO: Do this using hash map
	sort.Ints(nums1)
	sort.Ints(nums2)

	p1 := 0
	p2 := 0
	result := make([]int, 0)

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1++
			continue
		}

		if nums1[p1] > nums2[p2] {
			p2++
			continue
		}

		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		}
	}

	return result
}

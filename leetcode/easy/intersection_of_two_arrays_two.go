package easy

import "sort"

/*
Number: 350
Difficult: Easy
Link: https://leetcode.com/problems/intersection-of-two-arrays-ii/
Tags: Array, Hash Table, Two Pointers, Binary Search, Sorting
Status: Reviewed
*/
func intersect(nums1 []int, nums2 []int) []int {
	// Sorting solution
	// Time: O(M*log(M) + N*log(N))
	// Space: O(min(M, N))

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

func intersectHashMap(nums1 []int, nums2 []int) []int {
	// HashMap solution
	// Time: O(M+N)
	// Space: O(M)
	// Obs: If check the smallest array the space complexity will be O(min(N, M))

	mapping := make(map[int]int)
	var result []int

	for _, v := range nums1 {
		if _, ok := mapping[v]; ok {
			mapping[v]++
		} else {
			mapping[v] = 1
		}
	}

	for _, v := range nums2 {
		if value, ok := mapping[v]; ok {
			if value > 0 {
				mapping[v]--
				result = append(result, v)
			}
		}
	}

	return result
}

package easy

/*
Number: 88
Difficult: Easy
Link: https://leetcode.com/problems/merge-sorted-array/
Tags: Array, Two Pointers, Sorting
Status: Reviewed
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Time: O(M+N)
	// Space: O(M) -> Because I need to create a copy of the first array
	temp := make([]int, m)
	// Copy nums1 because it will store the merge array
	copy(temp, nums1)
	p1 := 0
	p2 := 0

	for i := 0; i < m+n; i++ {
		// If nums2 is bigger than n, add the next number of temp array
		// If temp is bigger than m, add the next number of nums2 array
		// Add the smallest number between temp[index] and nums2[index]
		if p2 >= n || (p1 < m && temp[p1] < nums2[p2]) {
			nums1[i] = temp[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}

// O(1) spacy complexity approach
func mergeBetterSpaceComplexity(nums1 []int, m int, nums2 []int, n int) {
	// If you start from the end and compare what number is bigger,
	// you won't need a nums1 copy.

	// Time: O(M + N)
	// Space: O(1)
	p1 := m - 1
	p2 := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if p2 < 0 || (p1 >= 0 && nums1[p1] > nums2[p2]) {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}

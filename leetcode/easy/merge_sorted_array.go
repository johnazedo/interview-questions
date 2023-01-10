package easy

/*
Difficult: Easy
Link: https://leetcode.com/problems/merge-sorted-array/
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1)
	p1 := 0
	p2 := 0

	for i := 0; i < m+n; i++ {
		if p2 >= n || (p1 < m && temp[p1] < nums2[p2]) {
			nums1[i] = temp[p1]
			p1++
		} else {
			nums1[i] = nums2[p2]
			p2++
		}
	}
}

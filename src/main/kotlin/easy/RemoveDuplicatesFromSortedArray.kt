package easy

object RemoveDuplicatesFromSortedArray {
    /*
    * Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
    * Difficult: Easy
    */
    fun removeDuplicates(nums: IntArray): Int {
        var p0 = 0;
        var p1 = 1;
        var count = 1;

        while(p1 < nums.size) {
            if(nums[p0] < nums[p1]) {
                if (p1 - p0 == 1 ){
                    count++
                    p0++
                    p1++
                } else {
                    p0++
                    nums[p0] = nums[p1]
                    count++
                }
                continue
            }
            if(nums[p0] == nums[p1]) p1++
        }
        return count
    }
}
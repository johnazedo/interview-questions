package easy

object RemoveElements {
    /*
    * Link: https://leetcode.com/problems/remove-element/
    * Difficult: Easy
    */

    fun removeElement(nums: IntArray, x: Int): Int {
        var insertPointer = 0

        for(i in nums.indices) {
            if(nums[i] != x) {
                nums[insertPointer] = nums[i]
                insertPointer++
            }
        }

        return insertPointer
    }
}
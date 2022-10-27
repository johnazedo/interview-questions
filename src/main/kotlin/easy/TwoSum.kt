package easy

class TwoSum {
    /*
    * Link: https://leetcode.com/problems/two-sum/
    * Difficulty: Easy
    */
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val hashMap = mutableMapOf<Int, Int>()

        // This is a linear solution with O(N) space complexity
        nums.forEachIndexed { index, value ->
            val diff = target - value

            if(hashMap.contains(diff)){
                return@twoSum intArrayOf(hashMap[diff]!!, index)
            }

            hashMap[value] = index
        }
        throw Throwable("This problem has no solution!")
    }
}
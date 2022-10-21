class TwoSum {
    /*
    * Link: https://leetcode.com/problems/two-sum/
    * Difficulty: Easy
    */
    fun twoSum(nums: IntArray, target: Int): IntArray {
        nums.forEachIndexed { i, one ->
            nums.forEachIndexed { j, two ->
                if((one + two) == target && i != j) {
                    return arrayOf(i, j).toIntArray()
                }
            }
        }
        throw Throwable("This problem has no solution!")
    }

}
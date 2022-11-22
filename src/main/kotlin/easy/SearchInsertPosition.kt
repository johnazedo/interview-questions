package easy

object SearchInsertPosition {
    fun searchInsert(nums: IntArray, target: Int): Int {
        var end = nums.size - 1
        var start = 0
        var pivot = 0

        while(start <= end) {
            pivot = start + (end - start)/2
            if(nums[pivot] == target) return pivot
            if(nums[pivot] > target) end = pivot - 1
            if(nums[pivot] < target) start = pivot + 1
        }

        return start
    }
}
package easy

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class RemoveDuplicatesFromSortedArrayTest {

    @ParameterizedTest(name = "given a array of ints {0}, then it should return the sorted non-duplicates {1} and the amount of numbers {2}")
    @MethodSource("arguments")
    fun execute(
            nums: IntArray,
            expected: Int,
            expectedNums: IntArray,
    ) {
        val actual = RemoveDuplicatesFromSortedArray.removeDuplicates(nums)
        assertEquals(expected, actual)
        for(i in 0 until actual) {
            assertEquals(expectedNums[i], nums[i])
        }
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
                Arguments.of(intArrayOf(1, 1, 2), 2, intArrayOf(1, 2, 2)),
                Arguments.of(intArrayOf(0,0,1,1,1,2,2,3,3,4), 5, intArrayOf(0, 1, 2, 3, 4, 2, 2, 3, 3, 4)),
                Arguments.of(intArrayOf(0, 1, 2, 2, 2, 3, 3, 4), 5, intArrayOf(0, 1, 2, 3, 4, 3, 3, 4)),
                Arguments.of(intArrayOf(1, 2), 2, intArrayOf(1, 2)),
                Arguments.of(intArrayOf(1), 1, intArrayOf(1))
        )
    }
}
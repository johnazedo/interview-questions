package easy

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals
import kotlin.test.assertNotEquals

class RemoveElementsTest {
    @ParameterizedTest(name = "given a array of ints {0} and a number {1}, then it should return a {2} and the amount of numbers {3}")
    @MethodSource("arguments")
    fun execute(
            nums: IntArray,
            value: Int,
            expectedNums: IntArray,
            expectedNumsSize: Int
    ) {
        val actual = RemoveElements.removeElement(nums, value)
        assertEquals(expectedNumsSize, actual)
        for(i in 0 until actual) {
            assertNotEquals(expectedNums[i], value)
        }
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
                Arguments.of(intArrayOf(3, 2, 2, 3), 3, intArrayOf(2, 2, 2, 3), 2),
                Arguments.of(intArrayOf(0, 1, 2, 2, 3, 0, 4, 2), 2, intArrayOf(0, 1, 4, 0, 3, 0, 4, 2), 5),
                Arguments.of(intArrayOf(2), 3, intArrayOf(2), 1),
                Arguments.of(intArrayOf(1), 1, intArrayOf(), 0),
                Arguments.of(intArrayOf(3, 3), 5, intArrayOf(3, 3), 2)
        )
    }
}
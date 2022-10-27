package easy

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class TwoSumTest {

    @ParameterizedTest(name = "given {0} and target {1}, then it should return {2}")
    @MethodSource("twoSumArguments")
    fun `given a array and a target, return the index of numbers`(
        array: IntArray,
        target: Int,
        expected: IntArray
    ) {
        val actual = TwoSum().twoSum(array, target)
        assertEquals(expected.asList(), actual.asList())
    }

    private companion object {
        @JvmStatic
        fun twoSumArguments(): Stream<Arguments> = Stream.of(
            Arguments.of(intArrayOf(3, 2, 4), 6, intArrayOf(1, 2)),
            Arguments.of(intArrayOf(3, 3), 6, intArrayOf(0,1)),
            Arguments.of(intArrayOf(2, 7, 11, 15), 9, intArrayOf(0,1)),
        )
    }
}
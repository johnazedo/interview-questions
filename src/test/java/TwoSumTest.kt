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
            Arguments.of(arrayOf(3, 2, 4).toIntArray(), 6, arrayOf(1, 2).toIntArray()),
            Arguments.of(arrayOf(3, 3).toIntArray(), 6, arrayOf(0,1).toIntArray()),
            Arguments.of(arrayOf(2, 7, 11, 15).toIntArray(), 9, arrayOf(0,1).toIntArray()),
        )
    }
}
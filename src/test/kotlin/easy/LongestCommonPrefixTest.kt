package easy

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class LongestCommonPrefixTest {
    @ParameterizedTest(name = "given a array of strings {0}, then it should return the longest common prefix {1}")
    @MethodSource("arguments")
    fun execute(
        strs: Array<String>,
        expected: String
    ) {
        val actual = LongestCommonPrefix().longestCommonPrefix(strs)
        assertEquals(expected, actual)
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
            Arguments.of(arrayOf("flower", "flow", "flight"), "fl"),
            Arguments.of(arrayOf("dog", "racecar", "car"), ""),
            Arguments.of(arrayOf("aaa", "aa", "aaa"), "aa"),
            Arguments.of(arrayOf("c", "acc", "ccc"), "")
        )
    }
}
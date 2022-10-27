package easy

import easy.PalindromeNumber
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class PalindromeNumberTest {
    @ParameterizedTest(name = "given a number {0}, then it should return {1}")
    @MethodSource("arguments")
    fun execute(
        number: Int,
        expected: Boolean
    ) {
        val actual = PalindromeNumber().isPalindrome(number)
        assertEquals(expected, actual)
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
            Arguments.of(121, true),
            Arguments.of(-121, false),
            Arguments.of(10, false),
        )
    }
}
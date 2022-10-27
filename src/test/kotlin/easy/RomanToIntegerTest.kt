package easy

import easy.RomanToInteger
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class RomanToIntegerTest {
    @ParameterizedTest(name = "given a string {0}, then it should return the integer {1}")
    @MethodSource("arguments")
    fun execute(
        s: String,
        expected: Int
    ) {
        val actual = RomanToInteger().romanToInt(s)
        assertEquals(expected, actual)
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
            Arguments.of("III", 3),
            Arguments.of("MCMXCIV", 1994),
            Arguments.of("LVIII", 58),
        )
    }
}
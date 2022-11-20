package easy

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class ValidParenthesesTest {

    @ParameterizedTest(name = "given a string {0}, then it should return the a boolean {1}")
    @MethodSource("arguments")
    fun execute(
            str: String,
            expected: Boolean
    ) {
        val actual = ValidParentheses.isValid(str)
        assertEquals(expected, actual)
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
                Arguments.of("()", true),
                Arguments.of("()[]{}", true),
                Arguments.of("(]", false),
                Arguments.of("([)]", false),
                Arguments.of("{[]}", true),
                Arguments.of("{[]}]]]", false),
                Arguments.of("[", false),
                Arguments.of("]", false)
        )
    }
}

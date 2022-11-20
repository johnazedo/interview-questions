package easy

object ValidParentheses {
    /*
   * Link: https://leetcode.com/problems/valid-parentheses/
   * Difficulty: Easy
   */

    private val charMap = mapOf(
            Pair('{', 1),
            Pair('}', -1),
            Pair('[', 2),
            Pair(']', -2),
            Pair('(', 3),
            Pair(')', -3)
    )


    fun isValid(s: String): Boolean {
        val stack = ArrayDeque<Int>()

        for(c in s) {
            val number = charMap[c]!!
            if(number > 0) {
                // Quando é um sinal de abrir
                stack.addLast(number)
            } else {
                if(stack.isEmpty()) return false
                // Quando é um sinal de fechar
                if (stack.last() + number == 0) {
                    stack.removeLast()
                } else return false
            }
        }
        return stack.isEmpty()
    }
}
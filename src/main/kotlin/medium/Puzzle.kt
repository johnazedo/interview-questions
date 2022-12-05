package medium

class Save {
    /*

    Given a two dimensional array of letters, find a given word written in any of the 8 directions.
    I.e.

    EXAMPLE
    Input: UBER

    Input:
    A  U  I  K  F  W  N
    W  Q  B  O  L  X  P
    T  L  A  E  R  E  S
    Y  Z  X  E  R  L  W

    Output: true

    */

    fun main(args: Array<String>) {
        val matrix = arrayOf(
                charArrayOf('A', 'U', 'I', 'K', 'F', 'W', 'N'),
                charArrayOf('W', 'Q', 'B', 'O', 'L', 'X', 'P'),
                charArrayOf('T', 'L', 'A', 'E', 'R', 'E', 'S'),
                charArrayOf('Y', 'Z', 'X', 'E', 'R', 'L', 'W')
        )
        val word = "UBER"
        var count = 0;
    }
}
package easy

class RomanToInteger {

    private val numberMap = mapOf<Char, Int>(
        Pair('I', 1),
        Pair('V', 5),
        Pair('X', 10),
        Pair('L', 50),
        Pair('C', 100),
        Pair('D', 500),
        Pair('M', 1000)
    )

    fun romanToInt(s: String): Int {
        // Vamos percorrendo o array de char e vendo se o próximo for
        // maior que o que está então é porque equivale a uma subtração
        var sum = 0;
        var i = 0;

        // O(1) -> É constante porque o conjunto dos números romanos é finito
        while(i < s.length) {
            val actual = numberMap[s[i]]!! // O(1)
            val hasNext = i < s.length - 1

            sum += if(hasNext) {
                val next = numberMap[s[i+1]]!! // O(1)
                val nextIsBigger = actual < next
                if(nextIsBigger) {
                    i += 1
                    next-actual
                } else actual
            } else actual
            i += 1
        }

        return sum
    }


}
package easy

class LongestCommonPrefix {
    /*
     * Link: https://leetcode.com/problems/longest-common-prefix/
     * Difficult: Easy
     */

    fun longestCommonPrefix(strs: Array<String>): String {

        // 0 <= i <= 200
        var i = 0;
        var prefix = "";
        var sameLetter = true

        // Esse while vai de 0 até o tamanho da primeira palavra
        // Aumentar o tamanho da palavra não aumenta a quantidade, logo não é O(Nˆ2)
        // O(1)
        while(i < strs[0].length && sameLetter) {
            val letter = strs[0][i]

            // Esse for percorre todas as strings no array
            // O(N)
            for(str in strs) {
                if(i<str.length) {
                    sameLetter = str[i] == letter
                    if(!sameLetter) break
                    continue
                }
                sameLetter = false
                break
            }

            if(sameLetter){
                prefix += letter
            }

            i++
        }

        return prefix

    }
}
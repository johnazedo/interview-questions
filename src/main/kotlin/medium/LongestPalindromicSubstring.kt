package medium

class LongestPalindromicSubstring {

    /*
    * Link: https://leetcode.com/problems/longest-palindromic-substring/
    * Difficult: Medium
    */
    fun longestPalindrome(s: String): String {
        return firstSolution(s)
    }

    private fun dynamicSolution(s: String): String {
        return ""
    }


    // TIME LIMIT EXCEEDED: 384ms
    private fun firstSolution(s: String): String {
        val sList = mutableListOf<String>()

        // Create a list with all palindrome
        for(i in s.indices) {
            for(j in i+1.. s.length) {
                val sub = s.substring(i, j)
                if(sub == sub.reversed())
                    sList.add(sub)
            }
        }

        // Get the biggest palindrome from list
        var max = 0
        var result: String = ""
        sList.forEach {
            if(it.length>max) {
                result = it
                max = it.length
            }
        }

        return result
    }
}
package easy

class PalindromeNumber {
    /*
     * Link: https://leetcode.com/problems/palindrome-number/
     * Difficult: Easy
     */
    fun isPalindrome(x: Int): Boolean {
        var reversedNumber: Int = 0
        var number = x

        if(x<0 || x % 10 == 0 && x != 0){
            return false
        }

        while (number > reversedNumber) {
            reversedNumber = reversedNumber * 10 + number % 10
            number /= 10
        }

        return reversedNumber == number || number == reversedNumber/10
    }
}
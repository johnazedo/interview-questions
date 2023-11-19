#include "../data_structures.h"
#include <bits/stdc++.h>

/*
 * Number: 09
 * Difficult: Easy
 * Link: https://leetcode.com/problems/palindrome-number
 * Tags: Math
 */ 

class PalindromeNumber {
public:
    /*
     * Time complexity: O(N)
     * Space complexity: O(N)
     * 
     * Time beats 14.16%
     * Memory beats 39.25%
     */
    bool isPalindrome(int x) {
        std::string s = std::to_string(x);
        int left = 0;
        int right = s.size() - 1;

        for(;;) {
            if(left >= right) {
                break; 
            }

            if(s[left] != s[right]) {
                return false;
            }

            left++;
            right--;
        }

        return true;
    }

    // Improve space complexity to O(1)
};
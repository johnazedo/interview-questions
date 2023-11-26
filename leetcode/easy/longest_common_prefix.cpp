#include "../data_structures.h"
#include <bits/stdc++.h>

using namespace std;

/*
 * Number: 14
 * Difficult: Easy
 * Link: https://leetcode.com/problems/longest-common-prefix/
 * Tags: String, Trie
 */ 

class LongestCommonPrefix {
public:
    /*
     * Time complexity: O(N+M)
     * Space complexity: O(M)
     * 
     * N is the size of array
     * M is the size of prefix
     * 
     * Time beats 5.78%
     * Memory beats 6.8%
     */
    string longestCommonPrefix(vector<string>& strs) {
        string aws = "";
        bool finish = false;

        for(int i = 0; i < strs[0].size(); i++) {
            for(string s: strs) {
                if(i >= s.size()) {
                    finish = true;
                    break;
                }

                if(strs[0][i] != s[i]) {
                    finish = true;
                    break;
                }      
            }

            if(finish) {
                break;
            }

            aws += strs[0][i];        
        }

        return aws;
    }
};
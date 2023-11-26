#include "../data_structures.h"
#include <bits/stdc++.h>

using namespace std;

/*
 * Number: 13
 * Difficult: Easy
 * Link: https://leetcode.com/problems/roman-to-integer/
 * Tags: Hash Table, Math, String
 */ 

class RomanToInteger {
public:
    /*
     * Time complexity: O(N)
     * Space complexity: O(1)
     * 
     * Time beats 39.71%%
     * Memory beats 62.11%
     */
    int romanToInt(string s) {
        map<char, int> dict;
        map<char, int>::iterator it_actual, it_next;

        dict['I'] = 1;
        dict['V'] = 5;
        dict['X'] = 10;
        dict['L'] = 50;
        dict['C'] = 100;
        dict['D'] = 500;
        dict['M'] = 1000;

        int aws = 0;

        for(int i = 0; i < s.size(); i++) {            
            it_actual = dict.find(s[i]);
            it_next = dict.end();
            
            if(i+1 < s.size()) {
                it_next = dict.find(s[i+1]);
            }
            
            if(it_next == dict.end()) {
                aws += it_actual->second;
                continue;
            }

            if(it_actual->second >= it_next->second) {
                aws += it_actual->second;
            } else {
                aws -= it_actual->second;
            }
        }

        return aws;
    }
};
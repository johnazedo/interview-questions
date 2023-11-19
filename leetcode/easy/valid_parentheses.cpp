#include "../data_structures.h"
#include <bits/stdc++.h>

using namespace std;

/*
 * Number: 20
 * Difficult: Easy
 * Link: https://leetcode.com/problems/valid-parentheses
 * Tags: String, Stack
 */ 

class ValidParentheses {
public:
    /*
     * Time complexity: O(N)
     * Space complexity: O(N)
     *
     * Time beats 45.79% 
     * Memory Beats 64.86%
     */
    bool isValid(string s) {
        vector<char> stack;
        map<char, char> dict;
        map<char, char>::iterator it;
        int index;

        dict[')'] = '(';
        dict[']'] = '[';
        dict['}'] = '{';

        for(char &c : s) {
            if(c == '(' || c == '[' || c == '{') {
                stack.push_back(c);
                continue;
            }

            if(stack.empty()) {
                return false;
            }
            
            it = dict.find(c);
            index = stack.size() - 1;
            if(it->second != stack[index]) {
                return false;
            }

            stack.pop_back();
        }

        return stack.empty();
    }
};

// TODO: Improve runtime performance to beats more than 50% of solutions

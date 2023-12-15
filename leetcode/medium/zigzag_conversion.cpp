#include "../data_structures.h"
#include <bits/stdc++.h>

using namespace std;

/*
 * Number: 06
 * Difficult: Medium
 * Link: https://leetcode.com/problems/zigzag-conversion/
 * Tags: 
 */ 

class Solution {
public:
    string convert(string s, int numRows) {
        if(s.size() <= numRows) {
            return s;
        }

        int cols = int(ceil(s.size()/(numRows-1)));
        vector<int> matrix[numRows];

        int line = 0;
        for(char& c: s) {}
    }
};
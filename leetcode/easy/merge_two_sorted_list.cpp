#include "../data_structures.h"
#include <bits/stdc++.h>

/*
 * Number: 21
 * Difficult: Easy
 * Link: https://leetcode.com/problems/merge-two-sorted-lists
 * Tags: Linked List, Recursion
 */ 

class RecursiveMergeTwoSortedList {
public:
    /*
    * Recursive approach
    *
    * Time: O(M+N)
    * Space: O(M+N) 
    * 
    * Time beats 85.5%
    * Memory beats 81.1%
    */
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        if(list1 == NULL) return list2;
        if(list2 == NULL) return list1;
        
        if(list1->val <= list2->val) {
            list1->next = mergeTwoLists(list1->next, list2);
            return list1;
        } else {
            list2->next = mergeTwoLists(list1, list2->next);
            return list2;
        }
    }
};

// TODO: Make the iteratively approach

#include "../data_structures.h"
#include <bits/stdc++.h>

using namespace std;

/*
 * Number: 02
 * Difficult: Medium
 * Link: https://leetcode.com/problems/add-two-numbers/
 * Tags: Linked List, Math, Recursion
 */ 

class AddTwoNumbers {
public:
    /*
     * Time complexity: O(N)
     * Space complexity: O(N)
     * 
     * Time beats: 84.96%
     * Space beats: 95.5%
     */
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        return helper(l1, l2);
    }

    ListNode* helper(ListNode* l1, ListNode* l2, int has_one = 0) {
        int sum, number, up;

        if(l1 == NULL && l2 == NULL && has_one == 0) {
            return NULL;
        }

        if(l1 == NULL && l2 == NULL && has_one != 0) {
            l1 =  new ListNode(has_one, NULL);
            return l1;
        }

        if(l1 == NULL) {
            sum = l2->val + has_one;
            l1 = new ListNode(0, NULL); 
        }

        if(l2 == NULL) {
            sum = l1->val + has_one;
            l2 = new ListNode(0, NULL); 
        }

        if(l1 != NULL && l2 != NULL) {
            sum = l1->val + l2->val + has_one;
        }

        if(sum >= 10) {
            number = sum % 10;
            up = 1;
        } else {
            number = sum;
            up = 0;
        }

        l1->val = number;
        l1->next = helper(l1->next, l2->next, up);

        return l1;
    }
};
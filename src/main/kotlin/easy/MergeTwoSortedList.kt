package easy

import datastructures.ListNode

class MergeTwoSortedList {
    /*
    * Link: https://leetcode.com/problems/merge-two-sorted-lists/
    * Difficulty: Easy
    * */
    fun mergeTwoLists(list1: ListNode?, list2: ListNode?): ListNode? {
        // Como as duas já estão ordenadas, podemos percorrer o array e
        // comparar cada elemento e decidir qual o maior
        var ll1 = list1
        var ll2 = list2

        // Esse algorítmo tem complexidade 0(N) de tempo no qual N é a quantidade
        // de elementos de cada lista somados enquanto o uso de memória é constante
        val head = ListNode(0)
        var temp = head
        while (ll1 != null || ll2 != null) {
            if((ll1?.`val` ?: 999) <= (ll2?.`val` ?: 999)){
                temp.next = ll1
                ll1 = ll1?.next
            } else {
                temp.next = ll2
                ll2 = ll2?.next
            }
            temp = temp.next!!
        }

        return head.next
    }
}
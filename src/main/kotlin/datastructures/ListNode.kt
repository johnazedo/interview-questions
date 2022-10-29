package datastructures


class ListNode(var `val`: Int) {
    var next: ListNode? = null

    fun putNext(node: ListNode): ListNode{
        this.next = node
        return this
    }
}
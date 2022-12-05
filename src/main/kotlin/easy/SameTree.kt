package easy

import datastructures.TreeNode

object SameTree {

    fun isSameTree(p: TreeNode?, q: TreeNode?): Boolean {
        return recursion(p, q)
    }

    fun recursion(p: TreeNode?, q: TreeNode?): Boolean {
        if(p == null && q == null) return true
        if(p == null || q == null) return false

        if(p.`val` != q.`val`) return false

        return recursion(p.left, q.left) && recursion(p.right, q.right)
    }
}
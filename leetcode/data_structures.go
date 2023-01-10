package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

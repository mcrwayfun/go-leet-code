package main

// time complexity: O(n)
// space complexity: O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}
	return p == nil && q == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

// time complexity: O(n)
// space complexity: O(n)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && helper(p.Left, q.Right) &&
			helper(p.Right, q.Left)
	}
	return p == nil && q == nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

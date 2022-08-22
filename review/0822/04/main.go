package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTreeRecursive(root.Left, root.Right)
}

func isSymmetricTreeRecursive(p, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val && isSymmetricTreeRecursive(p.Left, q.Right) &&
			isSymmetricTreeRecursive(p.Right, q.Left)
	}

	return p == nil && q == nil
}

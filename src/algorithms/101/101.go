package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}

	return isSymmetricHelper(root,root)
}

func isSymmetricHelper(p *TreeNode,q *TreeNode)bool{
	if p == nil && q == nil{
		return true
	}

	if p == nil || q == nil{
		return false
	}

	if p.Val != q.Val{
		return false
	}

	return isSymmetricHelper(p.Left,q.Right) &&
		isSymmetricHelper(p.Right,q.Left)
}
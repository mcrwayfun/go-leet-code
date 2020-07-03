package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if isEqual(s, t){
		return true
	}

	if s != nil && isSubtree(s.Left, t){
		return true
	}

	if s != nil && isSubtree(s.Right, t){
		return true
	}

	return false
}

func isEqual(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	// 值不相等
	if a.Val != b.Val {
		return false
	}

	// 递归判断左右子树
	return isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}
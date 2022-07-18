package main

// time complexity: O(n*lgn)
// space complexity: O(n)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftValid := root.Left == nil || root.Val > maxNode(root.Left).Val
	rightValid := root.Right == nil || root.Val < minNode(root.Right).Val
	return leftValid && rightValid &&
		isValidBST(root.Left) && isValidBST(root.Right)
}

func maxNode(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func minNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// time complexity: O(n)
// space complexity: O(n)
func isValidBSTV2(root *TreeNode) bool {
	return isValidBSTBound(root, nil, nil)
}

func isValidBSTBound(root *TreeNode, lower, upper *TreeNode) bool {
	if root == nil {
		return true
	}

	if lower != nil && root.Val <= lower.Val {
		return false
	}

	if upper != nil && root.Val >= upper.Val {
		return false
	}

	return isValidBSTBound(root.Left, lower, root) &&
		isValidBSTBound(root.Right, root, upper)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
time complexity: O(n*lgn)
space complexity: O(n)
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftValid := root.Left == nil || root.Val > TreeMax(root.Left).Val
	rightValid := root.Right == nil || root.Val < TreeMin(root.Right).Val
	return leftValid && rightValid &&
		isValidBST(root.Left) && isValidBST(root.Right)
}

func TreeMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func TreeMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
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

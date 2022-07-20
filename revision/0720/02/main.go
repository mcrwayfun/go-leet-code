package main

// time complexity: O(n)
// space complexity: O(1)
func isValidBST(root *TreeNode) bool {
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

package main

// time complexity: O(n)
// space complexity: O(n)
func isValidBST(root *TreeNode) bool {
	return isValidBSTBound(root, nil, nil)
}

// upper root 当前最大的节点，对于左子树来说，就是root
// lower root 当前最小的节点，对于右子树来说，就是root
func isValidBSTBound(root *TreeNode, upper, lower *TreeNode) bool {
	if root == nil {
		return true
	}

	// 存在最大的节点，但是当前 root 比最大节点都大，返回false
	if upper != nil && root.Val >= upper.Val {
		return false
	}

	// 存在最小的节点，但是当前 root 比最小节点的都小，返回false
	if lower != nil && root.Val <= lower.Val {
		return false
	}

	return isValidBSTBound(root.Left, root, lower) &&
		isValidBSTBound(root.Right, upper, root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

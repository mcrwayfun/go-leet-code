package main

/**
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：
输入：root = [2,1,3]
输出：true

示例 2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func isValidBST(root *TreeNode) bool {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/
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

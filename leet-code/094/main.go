package main

/**
给定一个二叉树的根节点 root ，返回 它的 中序遍历 。

示例 1：
输入：root = [1,null,2,3]
输出：[1,3,2]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func inorderTraversal(root *TreeNode) []int{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
 */

// time complexity: O(n)
// space complexity: O(n)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalHelper(root, &res)
	return res
}

func inorderTraversalHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalHelper(root.Right, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

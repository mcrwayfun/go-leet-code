package main

/**
题目：给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：true

示例 2：
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

示例 3：
输入：root = []
输出：true

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func isBalanced(root *TreeNode) bool{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
 */

// time complexity: O(n*lgn)
// space complexity: O(n)
// 假设是一棵满二叉树，因为每一层都需要递归，从第一层开始
// 1: 第一层1个节点，时间复杂度为O(n)
// 2: 第二层2个节点，单个节点的访问O(n/2)*2=O(n)
// 3：第n层n个节点，时间复杂度为O(n)
// 树的高度为lgn，所以时间复杂度为n*lgn
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(getHeight(root.Left)-getHeight(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// time complexity: O(n)
// space complexity: O(1)
// 由下往上，每个节点最多遍历一次
func isBalancedV2(root *TreeNode) bool {
	return getHeightAndCheck(root) != -1
}

/**
1: 根节点则返回0
2：如果非平衡，则返回-1
3：否则返回当前节点的高度
 */
func getHeightAndCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeightAndCheck(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeightAndCheck(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

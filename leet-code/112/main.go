package main

/**
给你二叉树的根节点root 和一个表示目标和的整数targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。

示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。
示例 2：


输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。
示例 3：

输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func hasPathSum(root *TreeNode, targetSum int) bool {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(n)
// space complexity: O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}

// time complexity: O(n)
// space complexity: O(n)
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stack = append([]*TreeNode{}, root)
	var sumStack = append([]int{}, targetSum)
	var s *TreeNode
	var n int
	for len(stack) > 0 {
		s, stack = stack[len(stack)-1], stack[:len(stack)-1]
		n, sumStack = sumStack[len(sumStack)-1], sumStack[:len(sumStack)-1]

		if s.Left == nil && s.Right == nil && s.Val == n {
			return true
		}
		if s.Left != nil {
			stack = append(stack, s.Left)
			sumStack = append(sumStack, n-s.Val)
		}
		if s.Right != nil {
			stack = append(stack, s.Right)
			sumStack = append(sumStack, n-s.Val)
		}
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

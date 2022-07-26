package main

/**
给你二叉树的根节点 root ，返回它节点值的前序遍历。

示例 1：
输入：root = [1,null,2,3]
输出：[1,2,3]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [1]
输出：[1]

示例 4：
输入：root = [1,2]
输出：[1,2]

示例 5：
输入：root = [1,null,2]
输出：[1,2]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func preorderTraversal(root *TreeNode) []int {}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(n)
// space complexity: O(n)
// 前序遍历：根左右
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack = []*TreeNode{root}
	var s *TreeNode
	var ret []int
	for len(stack) > 0 {
		s, stack = stack[len(stack)-1], stack[:len(stack)-1]
		ret = append(ret, s.Val)
		if s.Right != nil {// 利用栈先进后出的特性，先将right压入栈中
			stack = append(stack, s.Right)
		}
		if s.Left != nil {
			stack = append(stack, s.Left)
		}
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

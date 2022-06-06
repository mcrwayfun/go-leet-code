package main

/**
给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[15,7],[9,20],[3]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func levelOrderBottom(root *TreeNode) [][]int {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
 */


// time complexity: O(n)
// space complexity: O(1)
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i:=0; i<size; i++ {
			s, queue = queue[0], queue[1:]
			level = append(level, s.Val)

			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		res = append(res, level)
	}

	for i:=0; i<len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
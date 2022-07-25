package main

/**
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历，
postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。

示例 1:
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:
输入：inorder = [-1], postorder = [-1]
输出：[-1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func buildTree(inorder []int, postorder []int) *TreeNode {}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

// time complexity: O(n)
// space complexity: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	var posIn = make(map[int]int, len(inorder))
	for i, v := range inorder {
		posIn[v] = i
	}

	return buildTreeHelper(postorder, 0, len(postorder)-1, 0, posIn)
}

func buildTreeHelper(postorder []int, postStart, postEnd int,
	inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	// 根节点
	var root = &TreeNode{Val: postorder[postEnd]}
	var rootIdx = inPos[postorder[postEnd]]
	var leftLen = rootIdx - inStart
	// 二叉树左子树区间[postStart, postStart+leftLen-1]
	// 二叉树右子树区间[postStart+leftLen, postEnd-1] 因为postEnd是根节点
	root.Left = buildTreeHelper(postorder, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildTreeHelper(postorder, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

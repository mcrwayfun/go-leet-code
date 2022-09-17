package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for index, v := range inorder {
		inPos[v] = index
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, inPos)
}

func buildTreeHelper(preorder []int, preStart, preEnd int,
	inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// 前序遍历起始坐标就是二叉树的根节点
	var root = &TreeNode{Val: preorder[preStart]}
	// 找到根节点在中序遍历中的下标
	var rootIdx = inPos[preorder[preStart]]
	// 二叉树左子树个数
	var leftLen = rootIdx - inStart
	// 构建二叉树的左子树
	root.Left = buildTreeHelper(preorder, preStart + 1, preStart + leftLen, inStart, inPos)
	// 构建二叉树的右子树
	root.Right = buildTreeHelper(preorder, preStart + leftLen + 1, preEnd, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

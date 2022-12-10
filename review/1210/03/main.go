package main

// 前序遍历, 按照中左右的方式遍历二叉树
// 中序遍历, 按照左中右的方式遍历二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for k, v := range inorder {
		inPos[v] = k
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, inPos)
}

// 现在前序遍历中获取到 根节点 的值
// 根据值获取到 根节点 在中序遍历中的下标，如此可以得到 左子树 和 右子树 的节点个数
// 接下来可利用 前序遍历，根据节点个数递归构造左右子树
func buildTreeHelper(preorder []int, preStart, preEnd int,
	inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	rootIndex := inPos[preorder[preStart]]
	leftLen := rootIndex - inStart
	// preStart+1 是应为首位是 根节点
	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildTreeHelper(preorder, preStart+leftLen+1, preEnd, rootIndex+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

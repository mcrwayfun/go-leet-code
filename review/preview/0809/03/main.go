package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for i, v := range inorder {
		inPos[v] = i
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, inPos)
}

func buildTreeHelper(preorder []int, preStart int, preEnd int,
	inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	var root = &TreeNode{Val: preorder[preStart]}
	var rootIdx = inPos[preorder[preStart]]
	var leftLen = rootIdx - inStart
	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildTreeHelper(preorder, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

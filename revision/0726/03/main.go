package main

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

	var root = &TreeNode{Val: postorder[postEnd]}
	var rootIdx = inPos[postorder[postEnd]]
	var leftLen = rootIdx - inStart

	root.Left = buildTreeHelper(postorder, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildTreeHelper(postorder, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

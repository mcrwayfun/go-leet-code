package main

// 中序遍历：左中右
// 后序遍历：左右中
// time complexity: O(n)
// space complexity: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for i, v := range inorder {
		inPos[v] = i
	}
	return buildTreeHelper(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildTreeHelper(postorder []int, postStart, postEnd int,
	inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	var root = &TreeNode{Val: postorder[postEnd]}
	var rootIdx = inPos[postorder[postEnd]]
	var leftLen = rootIdx - inStart // 中序遍历中左子树的个数
	root.Left = buildTreeHelper(postorder, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildTreeHelper(postorder, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

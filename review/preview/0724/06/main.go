package main

// time complexity: O(n)
// space complexity: O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	var postIn = make(map[int]int, len(inorder))
	for i, v := range inorder {
		postIn[v] = i
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, postIn)
}

func buildTreeHelper(preorder []int, preStart, preEnd int,
	postStart int, postIn map[int]int) *TreeNode {

	if preStart > preEnd {
		return nil
	}

	var root = &TreeNode{Val: preorder[preStart]}
	rootIdx := postIn[preorder[preStart]]
	leftLen := rootIdx - postStart

	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftLen, postStart, postIn)
	root.Right = buildTreeHelper(preorder, preStart+leftLen+1, preEnd, rootIdx+1, postIn)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

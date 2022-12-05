package main

// time complexity: O(n)
// space complexity: O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	var inorderPos = make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderPos[v] = i
	}
	return buildTreeHelper(postorder, 0, len(postorder)-1, 0, inorderPos)
}

// 中序遍历 左中右
// 后序遍历 左右中
// 通过后序遍历最后一个节点可以得到 根节点 的值
// 通过 根节点 的值可以在中序遍历中得到 左右子树的个数
func buildTreeHelper(postorder []int, postStart, postEnd int,
	inStart int, inorderPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}

	var root = &TreeNode{Val: postorder[postEnd]}
	// 找到 根节点 在 中序遍历中的下标
	rootIndex := inorderPos[postorder[postEnd]]
	// 左子树的个数
	leftLen := rootIndex - inStart
	// 递归构建左子树
	root.Left = buildTreeHelper(postorder, postStart, postStart+leftLen-1, inStart, inorderPos)
	// 递归构建右子树
	root.Right = buildTreeHelper(postorder, postStart+leftLen, postEnd-1, rootIndex+1, inorderPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for i, v := range inorder {
		inPos[v] = i
	}
	return helper(postorder, 0, len(postorder)-1, 0, inPos)
}

func helper(postorder []int, postStart, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	var root = &TreeNode{Val: postorder[postEnd]}
	var inOrderRootIndex = inPos[root.Val] // 中序遍历中根节点的下标
	// 通过中序遍历可以得到根节点左子树的节点数和右子树的节点数
	var leftNodeNum = inOrderRootIndex - inStart

	// 得到左子树的节点个数，可以开始利用 postorder 递归构造树的结构
	// 优先构造左子树，起始下标是 postStart, 结束下标是 postStart+leftNodeNum-1
	root.Left = helper(postorder, postStart, postStart+leftNodeNum-1, inStart, inPos)

	// 构造右子树，起始下标是 postStart+leftNodeNum, 结束下标是 postEnd-1, 最后一个节点是根节点
	// 注意中序遍历中 inStart 的起始坐标是 inOrderRootIndex+1
	root.Right = helper(postorder, postStart+leftNodeNum, postEnd-1, inOrderRootIndex+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

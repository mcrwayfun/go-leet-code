package main

/**
给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func buildTree(preorder []int, inorder []int) *TreeNode {}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

/**
前序遍历的顺序是：中左右（顺序），中序遍历的顺序是：左中右
所以有以下的特性：
1. 前序遍历的第一个元素是根节点。
2. 根据前序遍历的根节点去中序遍历中寻找，根节点的左边是左子树，右边
是右子树。
3. 根据中序遍历左边元素的个数和右边元素的个数，可以推断出前序遍历中左子树
和右子树元素的位置。

现有preorder=[1,2,4,8,16], inorder=[2,1,8,4,16]。
1. 二叉树的根节点是1
2. 二叉树的中序遍历左子树是2，右子树是8，4，16
3. 二叉树的前序遍历左子树是2，右子树是4，8，16

time complexity: O(n)
space complexity: O(n)
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	var inPos = make(map[int]int, len(inorder))
	for index, v := range inorder {
		inPos[v] = index
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, 0, inPos)
}

/**
per: 前序遍历数组
preStart: 前序遍历起始坐标
preEnd: 前序遍历结束坐标
inStart: 中序遍历开始坐标
inPos: 中序遍历map，key=中序遍历数组元素，value=中序遍历数组元素下标
*/
func buildTreeHelper(pre []int, preStart, preEnd int,
	inStart int, inPos map[int]int) *TreeNode {

	// 前序遍历起始坐标 超过 结束坐标，不需要构造二叉树，直接返回
	if preStart > preEnd {
		return nil
	}

	// 前序遍历中起始坐标就是当前节点的根节点
	var root = &TreeNode{Val: pre[preStart]}
	// 根节点在中序遍历中的下标
	var rootIdx = inPos[pre[preStart]]
	// 二叉树左子树的个数=rootIdx-inStart
	var leftLen = rootIdx - inStart
	// 构造二叉树的左子树
	root.Left = buildTreeHelper(pre, preStart+1, preStart+leftLen, inStart, inPos)
	// 构造二叉树的右子树
	root.Right = buildTreeHelper(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

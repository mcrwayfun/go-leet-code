package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(1)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 找到比val应该插入的节点
	cur := root
	var last *TreeNode
	for cur != nil {
		last = cur
		if cur.Val > val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	if last != nil {
		if last.Val > val {
			last.Left = &TreeNode{Val: val}
		} else {
			last.Right = &TreeNode{Val: val}
		}
	}

	return root
}

func main() {
	nums := []int{4, 2, 7, 1, 3}
	root := Ints2TreeNode(nums)
	val := 5
	bst := insertIntoBST(root, val)
	result := Tree2InOrder(bst)
	fmt.Println(result)

	nums = []int{40, 20, 60, 10, 30, 50, 70}
	root = Ints2TreeNode(nums)
	val = 25
	bst = insertIntoBST(root, val)
	result = Tree2InOrder(bst)
	fmt.Println(result)

	val = 5
	bst = insertIntoBST(nil, val)
	result = Tree2InOrder(bst)
	fmt.Println(result)
}

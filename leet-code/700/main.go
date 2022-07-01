package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(1)
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	return dfs(root, val)
}

func dfs(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	left := dfs(root.Left, val)
	if left != nil {
		return left
	}
	if root.Val == val {
		return root
	}
	right := dfs(root.Right, val)
	if right != nil {
		return right
	}
	return nil
}

func main() {
	nodes := []int{4, 2, 7, 1, 3}
	root := Ints2TreeNode(nodes)
	val := 2
	bst := searchBST(root, val)
	fmt.Println(Tree2InOrder(bst))
}

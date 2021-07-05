package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(h)
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		var size = len(queue)
		var level []int
		for i := 1; i <= size; i++ {
			cur := queue[0]
			level = append(level, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			queue = queue[1:]
		}

		result = append(result, level)
	}

	return result
}

// time complexity: O(n)
// space complexity: O(n)
func levelOrderV2(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	dfs(root, &result, 0)
	return result
}

func dfs(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*result) == level {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)
	for _, v := range []*TreeNode{root.Left, root.Right} {
		dfs(v, result, level+1)
	}
	return
}

func main() {
	node := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := Ints2TreeNode(node)
	order := levelOrder(root)
	fmt.Println(order)

	order = levelOrderV2(root)
	fmt.Println(order)
}

package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(1)
func levelOrderBottom(root *TreeNode) [][]int {
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

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return result
}

func main() {
	node := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := Ints2TreeNode(node)
	order := levelOrderBottom(root)
	fmt.Println(order)
}

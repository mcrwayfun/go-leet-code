package main

import (
	. "../tool/tree"
	"fmt"
)

// time complexity: O(n)
// space complexity: O(n)
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	var reverseLevel bool
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
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

		if reverseLevel {
			for i := 0; i < len(level)/2; i++ {
				level[i], level[len(level)-i-1] = level[len(level)-i-1], level[i]
			}
		}

		reverseLevel = !reverseLevel
		result = append(result, level)
	}
	return result
}

func main() {
	node := []int{3, 9, 20, NULL, NULL, 15, 7}
	root := Ints2TreeNode(node)
	order := zigzagLevelOrder(root)
	fmt.Println(order)

	node2 := []int{1, 2, 3, 4, NULL, NULL, 5}
	root = Ints2TreeNode(node2)
	order = zigzagLevelOrder(root)
	fmt.Println(order)
}

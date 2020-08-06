package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func int2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != -1 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != -1 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var s, s1 []string
	dfs(root, &s, &s1)

	var sum int
	for _, v := range s1 {
		num, _ := strconv.Atoi(v)
		sum += num
	}
	return sum
}

func dfs(root *TreeNode, s *[]string, s1 *[]string) {
	if root == nil {
		return
	}

	*s = append(*s, fmt.Sprintf("%d", root.Val))

	if root.Left == nil && root.Right == nil {
		*s1 = append(*s1, strings.Join(*s, ""))
	}

	for _, v := range []*TreeNode{root.Left, root.Right} {
		dfs(v, s, s1)
	}

	*s = (*s)[:len(*s)-1]
	return
}

func main() {
	root := int2TreeNode([]int{1, 2, 3})
	numbers := sumNumbers(root)
	fmt.Println(numbers)

	root = int2TreeNode([]int{4, 9, 0, 5, 1})
	numbers = sumNumbers(root)
	fmt.Println(numbers)

	root = int2TreeNode([]int{0, 2, 3})
	numbers = sumNumbers(root)
	fmt.Println(numbers)
}

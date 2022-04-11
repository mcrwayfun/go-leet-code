package main

import (
	"fmt"
	"sort"
)

/**
题目：Given an array of unsorted numbers and a target number,
find all unique quadruplets in it, whose sum is equal to the target number.

描述：有一个未排序的数组和 target number，找到能够满足 四数之和 == target 的组合。

举例：
Example 1:

Input: [4, 1, 2, -1, 1, -3], target=1
Output: [-3, -1, 1, 4], [-3, 1, 1, 2]
Explanation: Both the quadruplets add up to the target.
Example 2:

Input: [2, 0, -1, 1, -2, 2], target=2
Output: [-2, 0, 2, 2], [-1, 0, 1, 2]
Explanation: Both the quadruplets add up to the target.

模板：
func searchQuadruplets(arr []int, target int) [][]int
*/

/**
思路：这道题比求 三数之和 和 target 的比较多了一位，这道题可以转换为 n^3的算法

time complexity: O(n^3)
space complexity: O(n)
*/
func searchQuadruplets(arr []int, target int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	sort.Ints(arr)

	var quadruplets [][]int
	for i := 0; i < len(arr)-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		for j := i + 1; j < len(arr)-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}

			searchPairs(arr, target, i, j, &quadruplets)
		}
	}

	return quadruplets
}

func searchPairs(arr []int, target int, first,
	second int, quadruplets *[][]int) {
	left := second + 1
	right := len(arr) - 1

	for left < right {
		quadrupletSum := arr[first] + arr[second] + arr[left] + arr[right]
		if quadrupletSum == target {
			*quadruplets = append(*quadruplets, []int{arr[first], arr[second], arr[left], arr[right]})
			left++
			right--

			for left < right && arr[left] == arr[left-1] {
				left++
			}

			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if quadrupletSum < target {
			left++
		} else {
			right--
		}
	}
}

func main() {
	arr := []int{4, 1, 2, -1, 1, -3}
	target := 1
	fmt.Println(searchQuadruplets(arr, target))

	arr2 := []int{2, 0, -1, 1, -2, 2}
	target2 := 2
	fmt.Println(searchQuadruplets(arr2, target2))
}

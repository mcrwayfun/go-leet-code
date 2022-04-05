package main

import (
	"fmt"
	"sort"
)

/**
题目：Given an array of unsorted numbers, find all unique triplets in it that add up to zero.

解析：给定一个未排序的数组，找到所有唯一的 三个数加起来等于0的。

举例：
Example 1:

Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.
Example 2:

Input: [-5, 2, -1, -2, 3]
Output: [[-5, 2, 3], [-2, -1, 3]]
Explanation: There are two unique triplets whose sum is equal to zero.

模板：
func searchTriplets(arr []int) [][]int
*/

/**
思路：
1：先将数组排序，[-3, 0, 1, 2, -1, 1, -2] -> [-3, -2, -1, 0, 1, 1, 2]
2：遍历数组，假设当前下标为i，数组长度为k，从当前下标开始：
	2-1：从区间[i+1,k]开始，判断是否存在元素a和b，能够使得 -arr[i]=a+b 存在
	2-2: 因为是排序的，遍历区间[i+1,k] 可以使用双指针的方式，left从i+1开始，right从k开始：
		if leftNum + rightNum == target then return leftNum and rightNum
		else if leftNum + rightNum > target then right--
		else left--

time complexity: O(n^2)
space complexity: O(n)
*/
func searchTriplets(arr []int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	sort.Ints(arr) // 按照升序排序

	var triplets [][]int
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}

		searchPairs(arr, -arr[i], i+1, &triplets)
	}

	return triplets
}

func searchPairs(arr []int, targetSum int, left int, triplets *[][]int) {
	right := len(arr) - 1
	for left < right {
		curSum := arr[left] + arr[right]
		if curSum == targetSum {
			*triplets = append(*triplets, []int{-targetSum, arr[left], arr[right]})
			left++
			right--

			for left < right && arr[left] == arr[left-1] {
				left++ // skip the same ele
			}

			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if curSum > targetSum {
			right--
		} else {
			left++
		}
	}
}

func main() {
	arr := []int{-3, 0, 1, 2, -1, 1, -2}
	result := searchTriplets(arr)
	fmt.Println(result)

	arr2 := []int{-5, 2, -1, -2, 3}
	result2 := searchTriplets(arr2)
	fmt.Println(result2)
}

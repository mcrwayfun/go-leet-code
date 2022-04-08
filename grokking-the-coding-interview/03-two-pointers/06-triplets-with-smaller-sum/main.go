package main

import (
	"fmt"
	"sort"
)

/**
题目：Given an array arr of unsorted numbers and a target sum, count all triplets
in it such that arr[i] + arr[j] + arr[k] < target where i, j,
and k are three different indices. Write a function to return the count of such triplets.

描述：有一个未排序的数组和target sum，求所有满足 arr[i] + arr[j] + arr[k] < target 的下标，
写一个函数返回满足条件的 三数之和 的数量。

举例：
Example 1:

Input: [-1, 0, 2, 3], target=3
Output: 2
Explanation: There are two triplets whose sum is less than the target: [-1, 0, 3], [-1, 0, 2]
Example 2:

Input: [-1, 4, 2, 1, 3], target=5
Output: 4
Explanation: There are four triplets whose sum is less than the target:
   [-1, 1, 4], [-1, 1, 3], [-1, 1, 2], [-1, 2, 3]

模板：
func searchTriplets(arr []int, target int) int
*/

/**
思路：
这道题是求 三数之和 < target 的场景，可以转换为 固定arr[i]，然后判断：
1：arr[j] + arr[k] < target then count++
2：arr[j] + arr[k] >= target then k--
*/
func searchTriplets(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	var count int
	for i := 0; i < len(arr); i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			tripletSum := arr[i] + arr[left] + arr[right]
			if tripletSum < target {
				count += right - left
				left++
			} else {
				right--
			}
		}
	}

	return count
}

func main() {
	arr := []int{-1, 0, 2, 3}
	target := 3
	fmt.Println(searchTriplets(arr, target))
}

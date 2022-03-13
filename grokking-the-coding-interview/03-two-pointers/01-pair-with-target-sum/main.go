package main

import "fmt"

/**
题目：Given an array of sorted numbers and a target sum,
find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair)
such that they add up to the given target.

描述：给一个排序的数组和一个sum，求数组中的一对数，使它们的和 与 sum 相等，返回这对数的下标。

举例：
Example 1:

Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6
Example 2:

Input: [2, 5, 9, 11], target=11
Output: [0, 2]
Explanation: The numbers at index 0 and 2 add up to 11: 2+9=11

模板：
func search(arr []int, targetSum int) []int
*/

/**
解题思路1：使用双指针来解决
由于数组已经排好序，可以利用有序数组的特性，不需要定义map来存储访问过的元素。
1：定义左指针从下标为0开始，右指针从数组末端开始。
	1-1：if arr[start] + arr[end] > target then end-- 因为要向左都是比 arr[end]要小的数
	1-2：if arr[start] + arr[end] < target then start++
	1-3：if arr[start] + arr[end] then return []int{start, end}
*/
func search(arr []int, targetSum int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}

	var start int
	var end = len(arr) - 1

	for start < end {
		leftNum := arr[start]
		rightNum := arr[end]

		if leftNum+rightNum > targetSum {
			end--
		} else if leftNum+rightNum < targetSum {
			start++
		} else if leftNum+rightNum == targetSum {
			return []int{start, end}
		}
	}

	return []int{-1, -1}
}

func main() {
	arr := []int{1, 2, 3, 4, 6}
	target := 6

	arr2 := []int{2, 5, 9, 11}
	target2 := 11

	fmt.Println(search(arr, target))
	fmt.Println(search(arr2, target2))
}

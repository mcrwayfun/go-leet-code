package main

import (
	"fmt"
)

/**
题目：Given an array of positive numbers and a positive number ‘S’,
find the length of the smallest contiguous subarray whose sum is
greater than or equal to ‘S’. Return 0, if no such subarray exists.

解析：给定一个正数数组和正数S，求最小的连续子数组的长度，其总和大于或等于S，
如果不存这样的子数组则返回0

举例：
Example 1:

Input: [2, 1, 5, 2, 3, 2], S=7
Output: 2
Explanation: The smallest subarray with a sum great than or equal to '7' is [5, 2].
Example 2:

Input: [2, 1, 5, 2, 8], S=7
Output: 1
Explanation: The smallest subarray with a sum greater than or equal to '7' is [8].
Example 3:

Input: [3, 4, 1, 1, 6], S=8
Output: 3
Explanation: Smallest subarrays with a sum greater than or equal to '8' are [3, 4, 1] or [1, 1, 6].

模板：
func findMinSubArray(s int, arr []int) int
*/

/**
思路1：使用暴力破解法
1：使用变量sum表达当前子数组的和，变量length表示当前子数组的长度
2：使用双层循环

time complexity: O(n^2)
space complexity: O(1)
*/
func findMinSubArray(s int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	minLength := len(arr) + 1
	for i := 0; i < len(arr); i++ {
		var sum int
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum >= s {
				minLength = min(minLength, j-i+1)
				break // 求的是最短长度，满足条件可以跳出内层循环了
			}
		}
	}

	if minLength == len(arr) + 1 {
		return 0
	}

	return minLength
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// input: 2, 1, 5, 2, 3, 2, s: 7
	// output: 2
	arr := []int{2, 1, 5, 2, 3, 2}
	s := 7

	// input: 2, 1, 5, 2, 8, s: 7
	// output: 1
	arr2 := []int{2, 1, 5, 2, 8}
	s2 := 8

	fmt.Println("findMinSubArray=", findMinSubArray(s, arr))
	fmt.Println("findMinSubArray=", findMinSubArray(s2, arr2))
}

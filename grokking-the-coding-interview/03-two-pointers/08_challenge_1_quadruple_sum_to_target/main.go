package main

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
思路：这道题比求 三数之和 和 target 的比较多了一位，可以转变成固定前 n-2个值，
剩余的去
*/

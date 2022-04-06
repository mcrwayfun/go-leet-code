package main

import (
	"fmt"
	"math"
	"sort"
)

/**
题目：Given an array of unsorted numbers and a target number, find a triplet
in the array whose sum is as close to the target number as possible,
return the sum of the triplet. If there are more than one such triplet,
return the sum of the triplet with the smallest sum.

解析：给出一个未排序的数组和一个目标数，在数组中找到一个三联体
中的一个三联体，其总和尽可能地接近目标数。
返回该三联体的总和。如果有多于一个这样的三联体
返回总和最小的那个三联体的总和。

举例：
Example 1:

Input: [-2, 0, 1, 2], target=2
Output: 1
Explanation: The triplet [-2, 1, 2] has the closest sum to the target.
Example 2:

Input: [-3, -1, 1, 2], target=1
Output: 0
Explanation: The triplet [-3, 1, 2] has the closest sum to the target.
Example 3:

Input: [1, 0, 1, 1], target=100
Output: 3
Explanation: The triplet [1, 1, 1] has the closest sum to the target.

模板：
func searchTriplet(arr []int, targetSum int) int
*/

/**
解题思路：这道题有点 dis。这道题是求最接近 targetSum 的三数之和，其实可以转换为
abs(targetSum - 三数之和) 最小。假设用 smallestDiff 来表示这个最小值，targetDiff 来表示
(targetSum - 三数之和) 差值，则有：
1：targetDiff == 0, 两者相等，直接返回 targetSum
2：targetDiff > 0, 表示三数之和不够大，需要将某个数值增大
3：targetDiff < 0, 表示三数之和太大，需要将某个数值减小

所以这里转换为，可以用双指针求 三数之和 在数组中的最小值
1：将数组排序
2：在当前 arr[i]，left=i+1, right=len(arr-1), 求[left, right] 区间内三数之和最小值
3：targetDiff > 0, 则 left++, 因为需要一个更大的值
4：targetDiff < 0, 则 right--, 因为需要一个更小的值
5：if abs(targetDiff) < abs(smallestDiff), then smallestDiff = targetDiff
6：结果返回 targetSum - smallestDiff

time complexity: O(n*lgn + n^2) = O(n^2)
space complexity: O(n)
*/
func searchTriplet(arr []int, targetSum int) int {
	if len(arr) == 0 {
		return 0
	}

	sort.Ints(arr)

	var smallestDiff = math.MaxInt64
	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]

			if targetDiff == 0 {
				return targetSum
			}

			if abs(targetDiff) < abs(smallestDiff) {
				smallestDiff = targetDiff
			}

			if targetDiff > 0 {
				left++ // 需要更大的数
			} else {
				right--
			}
		}
	}

	return targetSum - smallestDiff
}

func main() {
	arr := []int{-2, 0, 1, 2}
	targetSum := 2

	fmt.Println(searchTriplet(arr, targetSum))

	arr2 := []int{-3, -1, 1, 2}
	targetSum2 := 1
	fmt.Println(searchTriplet(arr2, targetSum2))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package main

import (
	"fmt"
	"math"
)

/**
题目：Given a sorted array, create a new array containing squares of
all the number of the input array in the sorted order.

描述：给定一个排序数组，创建一个新的数组包含数组中元素的所有平方，新创建的数组也要
保持顺序。

举例：
Example 1:

Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]
Example 2:

Input: [-3, -1, 0, 1, 2]
Output: [0 1 1 4 9]

模板：
func makeSquares(arr []int) []int
*/

/**
解题思路1：
1：找到数组中的正数最小值，假设为minPositiveNum
2：定义left和right，从minPositiveNum开始向左和向右遍历：
	leftNum = arr[left] * arr[left], rightNum = arr[right] * arr[right]
	if leftNum > rightNum then newArr[index] = leftNum
	else newArr[index]=rightNum

time complexity: O(n)
space complexity: O(1)
*/
func makeSquares(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	var minPositiveNum = math.MaxInt64
	var minPositiveNumIndex int

	// 找到数组中的正数最小值的下标
	for index, v := range arr {
		if v >= 0 && minPositiveNum > v {
			minPositiveNum = v
			minPositiveNumIndex = index
		}
	}

	var newArr []int
	left := minPositiveNumIndex - 1
	right := minPositiveNumIndex + 1

	newArr = append(newArr, arr[minPositiveNumIndex])
	for left >= 0 && right < len(arr) {
		leftNum := arr[left] * arr[left]
		rightNum := arr[right] * arr[right]

		if leftNum > rightNum {
			newArr = append(newArr, rightNum)
			right++
		} else {
			newArr = append(newArr, leftNum)
			left--
		}
	}

	for left >= 0 {
		newArr = append(newArr, arr[left]*arr[left])
		left--
	}

	for right < len(arr) {
		newArr = append(newArr, arr[right]*arr[right])
		right++
	}

	return newArr
}

func makeSquares2(arr []int) []int {
	var highIndex = len(arr) - 1
	var left = 0
	var right = len(arr) - 1
	ret := make([]int, len(arr))

	for left <= right {
		leftNum := arr[left] * arr[left]
		rightNum := arr[right] * arr[right]

		if leftNum > rightNum {
			ret[highIndex] = leftNum
			highIndex--
			left++
		} else {
			ret[highIndex] = rightNum
			highIndex--
			right--
		}
	}

	return ret
}

func main() {
	arr := []int{-2, -1, 0, 2, 3}
	fmt.Println(makeSquares(arr))

	arr2 := []int{-3, -1, 0, 1, 2}
	fmt.Println(makeSquares(arr2))

	fmt.Println(makeSquares2(arr))

	fmt.Println(makeSquares2(arr2))
}

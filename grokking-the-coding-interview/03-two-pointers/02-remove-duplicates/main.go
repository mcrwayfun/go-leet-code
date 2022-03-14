package main

import "fmt"

/**
题目：Given an array of sorted numbers, remove all duplicates from it.
You should not use any extra space; after removing the duplicates in-place
return the new length of the array.

描述：给一个排序数组，移除所有重复的元素。不可以使用额外的空间，返回移除重复元素后的数组长度。

举例：
Example 1:

Input: [2, 3, 3, 3, 6, 9, 9]
Output: 4
Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].
Example 2:

Input: [2, 2, 2, 11]
Output: 2
Explanation: The first two elements after removing the duplicates will be [2, 11].

模板：
func remove(arr []int) int
*/

/**
解题思路：双指针
由于不允许使用额外的空间（否则可以使用map来存储统计ele出现的次数）
1：使用双指针，start指针指向当前可以被替换的下标，end指针则遍历arr。
每当我们遇到一个非重复的数字，就将其移动到最后一个非重复的数字旁边。
*/
func remove(arr []int) int {
	var nextNonDuplicate = 1
	for i := 1; i < len(arr); i++ {
		if arr[nextNonDuplicate-1] != arr[i] {
			arr[nextNonDuplicate] = arr[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	arr2 := []int{2, 2, 2, 11}

	fmt.Println(remove(arr))
	fmt.Println(remove(arr2))
}

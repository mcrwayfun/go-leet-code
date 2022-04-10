package main

import "fmt"

/**
题目：Given an array containing 0s, 1s and 2s, sort the array in-place.
You should treat numbers of the array as objects, hence,
we can’t count 0s, 1s, and 2s to recreate the array.

描述：输入的数组中包含了0、1、2三种元素，请将这个数组原地排序，不可以创建新数组

举例：
Example 1:

Input: [1, 0, 2, 1, 0]
Output: [0 0 1 1 2]
Example 2:

Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]

模板：
func sort(arr []int) []int
*/

/**
思路：这是一道排序的题目，但是不允许使用新的数组。如果使用快排等
方法，时间复杂度将达到 O(lgn*n)。

因为这道题 元素 <= 3，所以可以考虑使用双指针的方式，在O(n)的时间复杂度
内完成排序。
1：假设left=0 and right=len(arr)-1。left指向的是元素0开始的下标，而right指向的是元素2开始的下标。
我们将元素0全部都搬到数组的前面，而元素2全部搬到数组的后面，中间留下的自然就是元素1了。
2：if arr[i] == 0, then swap(arr, i, left) 将元素0搬到前面，i++ and left++
3：if arr[i] == 1, then i++
4: if arr[i] == 2, then swap(arr, i, right) 将元素2搬到后面，right-- (注意这里i并没有++，因为
和right交换后，原本arr[right]的元素可能是0，则需要将它继续往前搬。

time complexity: O(n)
space complexity: O(1)
*/
func sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	left := 0
	right := len(arr) - 1
	for i := 0; i <= right; {
		if arr[i] == 0 {
			swap(&arr, i, left)
			i++
			left++
		} else if arr[i] == 1 {
			i++
		} else { // arr[i] == 2
			swap(&arr, i, right)
			right--
		}
	}

	return arr
}

func swap(arr *[]int, x, y int) {
	(*arr)[x], (*arr)[y] = (*arr)[y], (*arr)[x]
}

func main() {
	arr := []int{1, 0, 2, 1, 0}
	fmt.Println(sort(arr))

	arr2 := []int{2, 2, 0, 1, 2, 0}
	fmt.Println(sort(arr2))
}

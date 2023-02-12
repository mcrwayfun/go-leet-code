package main

import "fmt"

/**
核心思想：不使用新的数组
1：快排等排序算法时间复杂度为 O(n*lgn)
2：由于数组中只有三种元素，可以使用O(n)的方法来达到排序的效果
	2-1：申明两个下标start和end，分别表示可以替换为0和替换为2的位置
	2-2：遍历数组，当前ele为0则与下标start的元素交换
	2-3：当前ele为2则与下标end的元素交换
	2-4：当前ele为1则不移动
*/
func sort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	var start = 0
	var end = len(arr) - 1
	for i := 0; i <= end; {
		if arr[i] == 0 {
			arr[i], arr[start] = arr[start], arr[i]
			start++
			i++
		} else if arr[i] == 1 {
			i++
		} else {
			arr[i], arr[end] = arr[end], arr[i]
			end--
			// 不执行i++，因为当前end对应的ele可能是0
		}
	}

	return arr
}

func main() {
	arr := []int{1, 0, 2, 1, 0}
	fmt.Println(sort(arr))

	arr2 := []int{2, 2, 0, 1, 2, 0}
	fmt.Println(sort(arr2))
}

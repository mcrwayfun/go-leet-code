package main

import "fmt"

// 给一个排序数组，移除所有重复的元素。不可以使用额外的空间，返回移除重复元素后的数组长度。
// 2, 3, 3, 3, 6, 9, 9
func remove(arr []int) int {
	var i, j int
	for i = 0; i < len(arr); i++ {
		if arr[i] != arr[j] {
			if i-j > 1 {
				arr[j+1] = arr[i]
			}
			j++
		}
	}
	return j + 1
}

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(remove(arr))
}

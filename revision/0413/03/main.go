package main

import "fmt"

// 维持一个当前及之前不会存在重复元素的下标，即当前可以被替换的下标
func remove(arr []int) int {
	var nextNonDuplicate = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[nextNonDuplicate-1] {
			arr[nextNonDuplicate] = arr[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

// 使用两个指针，i和j。j指向有效数组的最后一个位置
// 当i和j指向的值不一致的时候，将i的值添加到j的下一个位置
func remove2(arr []int) int {
	var j int
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[j] {
			if i-j > 1 {// 如果全部数组都没有重复，则不用每次都进行交换的操作，可以先判断是否相邻，不相邻才需要交换
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

	arr = []int{2, 2, 2, 11}
	fmt.Println(remove(arr))

	arr = []int{2}
	fmt.Println(remove(arr))

	arr = []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(remove2(arr))

	arr = []int{2, 2, 2, 11}
	fmt.Println(remove2(arr))

	arr = []int{2}
	fmt.Println(remove2(arr))
}

package main

import "fmt"

// 维护两个指针，其中指针i遍历数组arr，指正j指向有效数组的最后一个下标
// if arr[i] != arr[j] 则 j++ && arr[j]=arr[i]
func remove(arr []int) int {
	var i, j int
	for i = 0; i < len(arr); i++ {
		if arr[i] != arr[j] {
			if i-j > 1 { // 相邻元素不必交换
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
}

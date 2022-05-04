package main

import "fmt"

func sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var start int          // 0可以被替换的开始下标
	var end = len(arr) - 1 // 2可以被替换的下标
	for i := 0; i <= end; {
		num := arr[i]
		if num == 0 {
			arr[i], arr[start] = arr[start], arr[i]
			i++
			start++
		} else if num == 1 {
			i++ // 1不用替换，最后会自然排序到中间
		} else {
			arr[i], arr[end] = arr[end], arr[i]
			end-- // 注意这里i并不需要++，因为end指向的元素可能是0，意味着还需要交换
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

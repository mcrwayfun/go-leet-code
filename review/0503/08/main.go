package main

import "fmt"

func sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var start int          // 可以放置0的下标
	var end = len(arr) - 1 // 可以放置1的下标
	for i := 0; i <= end; {
		num := arr[i]
		if num == 0 {
			arr[i], arr[start] = arr[start], arr[i]
			i++
			start++
		} else if num == 1 {
			i++
		} else {
			arr[i], arr[end] = arr[end], arr[i]
			end--
			// i++ 这里不作用，因为end--后可能会是0
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

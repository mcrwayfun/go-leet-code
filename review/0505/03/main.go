package main

import "fmt"

func sort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	var start int
	var end = len(arr) - 1
	for i := 0; i <= end; {
		num := arr[i]
		if num == 0 {
			arr[i], arr[start] = arr[start], arr[i]
			i++
			start++
		} else if num == 1 {
			i++ // no need to swap
		} else if num == 2 {
			arr[i], arr[end] = arr[end], arr[i]
			end--
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

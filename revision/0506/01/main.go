package main

import "fmt"

func remove(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var nextNonDuplicate int
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[nextNonDuplicate] {
			nextNonDuplicate++
			arr[nextNonDuplicate] = arr[i]
		}
	}
	return nextNonDuplicate + 1
}

func main() {
	arr := []int{2, 3, 3, 3, 6, 9, 9}
	fmt.Println(remove(arr))
}

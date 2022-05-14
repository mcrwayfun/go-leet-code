package main

import "fmt"

func remove(arr []int) int {
	var nextNonDuplicate int
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[nextNonDuplicate] {
			nextNonDuplicate++
			arr[i], arr[nextNonDuplicate] = arr[nextNonDuplicate], arr[i]
		}
	}
	return nextNonDuplicate + 1
}

func main() {
	fmt.Println(remove([]int{2, 3, 3, 3, 6, 9, 9}))
	fmt.Println(remove([]int{2, 2, 2, 11}))
}

package main

import "fmt"

func findLength(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	var oneRepeatCount int
	var start int
	var maxLength int

	for end := 0; end < len(arr); end++ {
		if arr[end] == 1 {
			oneRepeatCount++
		}

		for end-start+1-oneRepeatCount > k {
			if arr[start] == 1 {
				oneRepeatCount--
			}
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func main() {
	arr := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	k := 2
	fmt.Println(findLength(arr, k))
}

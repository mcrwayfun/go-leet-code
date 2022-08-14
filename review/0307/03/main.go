package main

import "fmt"

func findLength(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var maxLength int
	var numOnesCount int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			numOnesCount++
		}

		if windowEnd-windowStart+1-numOnesCount > k {
			if arr[windowStart] == 1 {
				numOnesCount--
			}
			windowStart++
		}

		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	arr := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	k := 2

	arr2 := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1}
	k2 := 3

	fmt.Println(findLength(arr, k))
	fmt.Println(findLength(arr2, k2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

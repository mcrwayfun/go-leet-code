package main

import "fmt"

func findMinSubArray(s int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var windowSum int
	var minLength = len(arr) + 1

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		for windowSum >= s {
			minLength = min(minLength, windowEnd-windowStart+1)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	if minLength == len(arr)+1 {
		return 0
	}

	return minLength
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	arr := []int{2, 1, 5, 2, 3, 2}
	s := 7

	arr2 := []int{2, 1, 5, 2, 8}
	s2 := 7

	fmt.Println(findMinSubArray(s, arr))
	fmt.Println(findMinSubArray(s2, arr2))
}

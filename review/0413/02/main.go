package main

import "fmt"

func findMinSubArray(s int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var minLength = len(arr) + 1
	var windowSum int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		for windowSum >= s {
			minLength = min(minLength, windowEnd-windowStart+1)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	if minLength == len(arr) + 1 {
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
	fmt.Println(findMinSubArray(s, arr))

	arr = []int{2, 1, 5, 2, 8}
	s = 7
	fmt.Println(findMinSubArray(s, arr))

	arr = []int{3, 4, 1, 1, 6}
	s = 8
	fmt.Println(findMinSubArray(s, arr))
}
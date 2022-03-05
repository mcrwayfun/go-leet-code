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

func main() {
	arr := []int{2, 1, 5, 2, 3, 2}
	k := 7

	arr2 := []int{2, 1, 5, 2, 8}
	k2 := 7

	arr3 := []int{3, 4, 1, 1, 6}
	k3 := 8

	fmt.Println(findMinSubArray(k, arr))
	fmt.Println(findMinSubArray(k2, arr2))
	fmt.Println(findMinSubArray(k3, arr3))

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

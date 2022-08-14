package main

import "fmt"

func findMaxSumSubArray(k int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowSum int
	var maxSum int
	var windowStart int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd-windowStart+1 == k {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3

	arr2 := []int{2, 3, 4, 1, 5}
	k2 := 2

	fmt.Println(findMaxSumSubArray(k, arr))
	fmt.Println(findMaxSumSubArray(k2, arr2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

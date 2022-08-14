package main

import "fmt"

func findMaxSumSubArray(k int, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var windowStart int
	var windowSum int
	var maxSum int

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
	fmt.Println(findMaxSumSubArray(k, arr))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

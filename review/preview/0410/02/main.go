package main

import "fmt"

func findAverages(k int, arr []int) []float64 {
	if len(arr) == 0 {
		return []float64{}
	}

	var windowSum int
	var averages []float64
	var windowStart int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]

		if windowEnd-windowStart+1 == k {
			averages = append(averages, float64(windowSum)/float64(k))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return averages
}

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	fmt.Println(findAverages(k, arr))
}

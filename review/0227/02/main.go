package main

import "fmt"

// time complexity: O(n)
// space complexity: O(1)
func findAverages(k int, arr []int) []float64 {
	if len(arr) == 0 {
		return []float64{}
	}

	var windowStart int
	var windowSum int
	var ret []float64

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		if windowEnd-windowStart+1 == k {
			ret = append(ret, float64(windowSum)/float64(k))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return ret
}

func main() {
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	fmt.Println(findAverages(k, arr))
}

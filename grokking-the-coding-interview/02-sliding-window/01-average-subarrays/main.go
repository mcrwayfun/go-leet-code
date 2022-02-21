package main

import "fmt"

/**
题目：Given an array, find the average of all contiguous subarrays of size ‘K’ in it.
解析：在给定的数组中，找出其中所有大小为k的连续子数的平均值。

举例：
Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
1：For the first 5 numbers (subarray from index 0-4),
the average is: (1+3+2+6-1)/5 => 2.2(1+3+2+6−1)/5=>2.2

2：The average of next 5 numbers (subarray from index 1-5) is:
(3+2+6-1+4)/5 => 2.8(3+2+6−1+4)/5=>2.8

3：For the next 5 numbers (subarray from index 2-6),
the average is: (2+6-1+4+1)/5 => 2.4(2+6−1+4+1)/5=>2.4

题目模板：
func findAverages(k int, arr []int) []float64 {

}

*/

/**
解题思路1：使用暴力破解法，使用两层循环的方式来解决。
需要注意的是，外层循环i可以只计算到 len(arr)-k, 否则内层循环容易out of index
time complexity: O(n^2)
space complexity: O(1)
*/
func findAverages(k int, arr []int) []float64 {
	var ret []float64
	if len(arr) == 0 {
		return ret
	}

	// 只用计算到前 len(arr)-k个即可
	for i := 0; i <= len(arr)-k; i++ {
		var sum int
		for j := i; j < k+i; j++ {
			sum += arr[j]
		}
		ret = append(ret, float64(sum)/float64(k))
	}
	return ret
}

/**
解题思路2：滑动窗口，滑动窗口核心的逻辑是维护used元素，以避免遍历多次。
1：使用变量sum来存储当前子数组的和
2：同时维护另一个变量 windowStart，当 windowEnd-windowStart == k时，sum = sum - arr[windowStart]，
并且windowStart++
3：注意这里的外层循环需要 windowEnd < len(arr) 而不是 windowEnd < len(arr)-k !!

time complexity: O(n)
space complexity: O(1)
*/
func findAverages2(k int, arr []int) []float64 {
	var sum int
	var windowStart int
	var ret []float64
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		sum += arr[windowEnd]
		if windowEnd-windowStart+1 == k {
			ret = append(ret, float64(sum)/float64(k))
			sum -= arr[windowStart]
			windowStart++
		}
	}
	return ret
}

func main() {
	// input: 1, 3, 2, 6, -1, 4, 1, 8, 2
	// output: 2.2 2.8 2.4 3.6 2.8
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5
	result := findAverages(k, arr)
	fmt.Println(result)

	result = findAverages2(k, arr)
	fmt.Println(result)
}

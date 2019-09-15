package _52_peak_index_in_mountain_array

// time complexity:O(n)
// space complexity:O(1)
func peakIndexInMountainArray(A []int) int {
	for i := 0; i < len(A); i++ {
		if i == len(A)-1 || A[i] > A[i+1] {
			return i
		}
	}

	return -1
}

// binary search
// time complexity:O(lgn)
// space complexity:O(1)
func peakIndexInMountainArray2(A []int) int {
	left, mid, right := 0, 0, len(A)-1

	for left < right {
		mid = left + (right-left)/2
		if A[mid] < A[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

/**
我们把符合下列属性的数组 A 称作山脉：

A.length >= 3
存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。

链接：https://leetcode-cn.com/problems/peak-index-in-a-mountain-array
*/
func peakIndexInMountainArray3(A []int) int {
	if len(A) < 3 {
		return -1
	}

	start, end := 0, len(A)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] < A[mid+1] {
			start = mid
		} else {
			end = mid
		}
	}

	if A[start] > A[end] {
		return start
	} else {
		return end
	}
}

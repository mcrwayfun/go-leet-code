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

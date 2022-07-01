package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		return (findKth(nums1, 0, nums2, 0, n/2) +
			findKth(nums1, 0, nums2, 0, n/2+1)) / 2.0
	}
	return findKth(nums1, 0, nums2, 0, n/2+1)
}

// find kth number of two sorted array
func findKth(A []int, startA int, B []int, startB int, k int) float64 {
	if startA >= len(A) {
		return float64(B[startB+k-1]) // 返回B的第k个数
	}

	if startB >= len(B) {
		return float64(A[startA+k-1]) // 返回A的第k个数
	}

	if k == 1 {
		return float64(min(A[startA], B[startB]))
	}

	halfOfKthA := getHalfOfKth(startA, A, k)
	halfOfKthB := getHalfOfKth(startB, B, k)

	if halfOfKthA < halfOfKthB {
		return findKth(A, startA+k/2, B, startB, k-k/2)
	}
	return findKth(A, startA, B, startB+k/2, k-k/2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getHalfOfKth(index int, num []int, k int) int {
	if index+k/2-1 < len(num) {
		return num[index+k/2-1]
	}
	return 1<<31 - 1
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6}
	nums2 := []int{2, 3, 4, 5}
	ans := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ans)
}

package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[len(nums)-k]
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	mid := length / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left []int, right []int) []int {
	var ret []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		ret = append(ret, left...)
	}

	if len(right) != 0 {
		ret = append(ret, right...)
	}
	return ret
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}

	index := sort(arr, left, right)
	_quickSort(arr, left, index-1)
	_quickSort(arr, index+1, right)
	return arr
}

func sort(arr []int, left, right int) int {
	i, j := left, right
	index := left // 基准

	for i != j {
		for arr[j] >= arr[index] && i < j {
			j--
		}

		for arr[i] <= arr[index] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[index], arr[i] = arr[i], arr[index]
	return i
}

func main() {
	num := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	ans := findKthLargest(num, 5)
	fmt.Println(ans)
}

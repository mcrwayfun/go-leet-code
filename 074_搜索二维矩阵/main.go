package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	if matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	column := len(matrix[0])
	start, end := 0, row*column-1
	for start+1 < end {
		mid := start + (end-start)/2
		num := matrix[mid/column][mid%column]
		if num == target {
			return true
		} else if num < target {
			start = mid
		} else {
			end = mid
		}
	}

	if matrix[start/column][start%column] == target ||
		matrix[end/column][end%column] == target {
		return true
	}

	return false
}

func main() {
	nums := [][]int{
		{
			1, 3, 5, 7,
		},
		{
			10, 11, 16, 20,
		},
		{
			23, 30, 34, 60,
		},
	}

	exist := searchMatrix(nums, 11)
	fmt.Println(exist)
}

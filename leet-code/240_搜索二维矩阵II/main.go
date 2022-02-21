package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	if matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}

	row, column := len(matrix)-1, len(matrix[0])-1
	x, y := row, 0
	for x >= 0 && y <= column {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			y++
		} else {
			x--
		}
	}
	return false
}

func searchMatrixV2(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	if matrix[0] == nil || len(matrix[0]) == 0 {
		return false
	}

	row := len(matrix)
	column := len(matrix[0])
	x, y := 0, column-1
	for x <= row-1 && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func main() {
	nums := [][]int{
		{
			1, 4, 7, 11, 15,
		},
		{
			2, 5, 8, 12, 19,
		},
		{
			3, 6, 9, 16, 22,
		},
		{
			10, 13, 14, 17, 24,
		},
		{
			18, 21, 23, 26, 30,
		},
	}

	fmt.Println(searchMatrix(nums, 5))
	fmt.Println(searchMatrix(nums, 31))

	nums2 := [][]int{
		{
			1,
		},
		{
			1,
		},
	}

	fmt.Println(searchMatrixV2(nums2, 2))
}

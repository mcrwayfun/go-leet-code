package main

// time complexity: (n^2)
// space complexity: O(n^2)
func generate(numRows int) [][]int {
	var triangle = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

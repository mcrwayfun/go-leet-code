package main

// time complexity: O(n^2)
// space complexity: O(n^2)
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ { // i 表示行数，第i行有i+1个元素
		rows := make([]int, i+1)
		rows[0] = 1
		rows[i] = 1
		for j := 1; j < i; j++ {
			rows[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = rows
	}
	return res
}

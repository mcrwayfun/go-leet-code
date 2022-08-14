package main

// time complexity: O(n^2)
// space complexity: O(n^2)
func generate(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}

	var ret = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		var rows = make([]int, i+1) // i行有i+1个元素
		rows[0] = 1
		rows[i] = 1
		for j := 1; j < i; j++ {
			rows[j] = ret[i-1][j-1] + ret[i-1][j]
		}
		ret[i] = rows
	}
	return ret
}

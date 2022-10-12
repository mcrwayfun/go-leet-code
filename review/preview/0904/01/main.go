package main

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	var d = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		var rows = make([]int, i+1)
		rows[0] = 1
		rows[i] = 1
		for j := 1; j < i; j++ {
			rows[j] = d[i-1][j-1] + d[i-1][j]
		}
		d[i] = rows
	}
	return d
}

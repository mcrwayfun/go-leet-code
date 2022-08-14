package main

// time complexity: O(n^2)
// space complexity: O(n^2)
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	var res = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		var ans = make([]int, numRows)
		ans[0] = 1
		ans[i] = 1
		for j := i + 1; j < numRows; j++ {
			ans[j] = res[j-1][i-1] + res[j][i-1]
		}
		res = append(res, ans)
	}
	return res
}

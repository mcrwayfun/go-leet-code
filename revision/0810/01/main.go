package main

// time complexity: O(n)
// space complexity: O(1)
func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	var res = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		var ans = make([]int, i+1) // 每一行最多i+1个元素
		ans[0] = 1
		ans[i] = 1
		for j := 1; j < i; j++ {
			ans[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = ans
	}
	return res
}

package main

import (
	"fmt"
	"strings"
)

func solveNQueens(n int) [][]string {
	ans := make([][]string, n)
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, n)
	}
	backtrack(n, 0, res, &ans)
	return ans
}

func backtrack(n, row int, res [][]string, ans *[][]string) {
	if row == n {
		board := make([]string, len(res))
		for i := 0; i < n; i++ {
			board[i] = strings.Join(res[i], "")
		}
		*ans = append(*ans, board)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(row, col, res, n) {
			res[row][col] = "Q"
			backtrack(n, row+1, res, ans)
			res[row][col] = "."
		}
	}
	return
}

func isValid(row, col int, res [][]string, n int) bool {
	for i := 0; i < row; i++ {
		if res[i][col] == "Q" {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if res[i][j] == "Q" {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if res[i][j] == "Q" {
			return false
		}
	}
	return true
}

func main() {
	ans := solveNQueens(4)
	fmt.Println(ans)
}

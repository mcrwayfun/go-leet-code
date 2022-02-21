package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	if n == 0 {
		return ans
	}

	var track string
	backtrack(n, n, track, &ans)
	return ans
}

func backtrack(left, right int, track string, ans *[]string) {
	if right < left {
		return
	}

	if left < 0 || right < 0 {
		return
	}

	if left == 0 && right == 0 {
		*ans = append(*ans, track)
		return
	}

	track = track + "("
	backtrack(left-1, right, track, ans)
	track = track[:len(track)-1]

	track = track + ")"
	backtrack(left, right-1, track, ans)
	track = track[:len(track)-1]
}

func main() {
	parenthesisa := generateParenthesis(3)
	fmt.Println(parenthesisa)
}

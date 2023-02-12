package main

import "fmt"

// time complexity: O(n^2)
// space complexity: O(1)
func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	var visited = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		visited[i] = make([]bool, len(s))
	}

	var n = len(s)
	var count int
	for i := n - 1; i >= 0; i-- {
		for j := i; j <= n-1; j++ {
			if i == j {
				visited[i][j] = true
			} else if i+1 == j && s[i] == s[j] {
				visited[i][j] = true
			} else if s[i] == s[j] && visited[i+1][j-1] {
				visited[i][j] = true
			}

			if visited[i][j] {
				count++
			}
		}
	}

	return count
}

func countSubstringsV2(s string) int {
	if len(s) == 0 {
		return 0
	}

	var count int
	for i := 0; i < len(s); i++ {
		count += expand(s, i, i)
		count += expand(s, i, i+1)
	}
	return count
}

func expand(s string, left int, right int) int {
	var count int
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstringsV2("aaa"))
}

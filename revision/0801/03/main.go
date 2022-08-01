package main

// time complexity: O(2^n)
// space complexity: O(n^2)
func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	var d = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		d[i] = make([]bool, len(s))
	}

	var n = len(s)
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				d[i][j] = true
			} else if i+1 == j {
				d[i][j] = s[i] == s[j]
			} else {
				d[i][j] = s[i] == s[j] && d[i+1][j-1]
			}
		}
	}

	var res [][]string
	var ele []string
	helper(s, 0, d, &res, &ele)
	return res
}

func helper(s string, start int, d [][]bool, res *[][]string, ele *[]string) {
	if start >= len(s) {
		*res = append(*res, append([]string{}, *ele...))
	} else {
		for end := start; end < len(s); end++ {
			if d[start][end] {
				*ele = append(*ele, s[start:end+1])
				helper(s, end+1, d, res, ele)
				*ele = (*ele)[:len(*ele)-1]
			}
		}
	}
}

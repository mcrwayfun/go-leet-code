package main

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	var n = len(s)
	var d = make([][]bool, n)
	for i := range d {
		d[i] = make([]bool, n)
	}

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
	var ans []string
	helper(s, 0, &ans, &res, d)
	return res
}

func helper(s string, start int, ans *[]string, res *[][]string, d [][]bool) {
	if start >= len(s) {
		*res = append(*res, append([]string{}, *ans...))
	} else {
		for end := start; end < len(s); end++ {
			if d[start][end] {
				*ans = append(*ans, s[start:end+1])
				helper(s, end+1, ans, res, d)
				*ans = (*ans)[:len(*ans)-1]
			}
		}
	}
}

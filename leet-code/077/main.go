package main

import "fmt"

func combine(n int, k int) [][]int {
	var ans [][]int
	var res []int
	backtrack(1, n, k, res, &ans)
	return ans
}

func backtrack(start, n, k int, res []int, ans *[][]int) {
	if len(res) == k {
		*ans = append(*ans, append([]int{}, res...))
		return
	}

	for i := start; i <= n; i++ {
		if used(res, i) {
			continue
		}

		res = append(res, i)
		backtrack(i, n, k, res, ans)
		res = res[:len(res)-1]
	}
}

func used(res []int, target int) bool {
	for _, v := range res {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	ans := combine(4, 2)
	fmt.Println(ans)

	ans = combine(1, 1)
	fmt.Println(ans)

	ans = combine(4, 1)
	fmt.Println(ans)
}

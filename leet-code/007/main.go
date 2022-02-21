package main

// time complexity:O(n)
// space complexity:O(1)
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	flag := 1
	if x < 0 {
		flag = -1
		x = -x
	}

	var ans int64 // 防止溢出
	for x != 0 {
		ans = ans*10 + int64(x%10)
		x = x/10
	}

	ans = ans * int64(flag)
	if ans < -(1<<31) || ans > 1<<31-1 {
		return 0
	}

	return int(ans)
}

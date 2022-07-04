package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	var res = n
	for res % 5 == 0 {
		res = res/5
	}
	for res % 3 == 0 {
		res = res/3
	}
	for res % 2 == 0 {
		res = res/2
	}
	return res == 1
}

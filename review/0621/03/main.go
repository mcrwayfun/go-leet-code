package main

func isUgly(n int) bool {
	var res = n
	if res <= 0 {
		return false
	}

	for res % 2 == 0 {
		res /= 2
	}
	for res % 3 == 0 {
		res /= 3
	}
	for res % 5 == 0 {
		res /= 5
	}
	return res == 1
}

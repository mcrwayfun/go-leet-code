package main

// time complexity: O(lgn)
// space complexity: O(1)
func firstBadVersion(n int) int {
	start := 0
	end := n
	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid
		}
	}

	if isBadVersion(end) {
		return end
	}
	return start
}

func isBadVersion(version int) bool {
	return false
}

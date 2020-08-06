package _01

func TwoSum(nums []int, target int) []int {
	// nums is nil or empty
	if nums == nil || len(nums) == 0 {
		return nil
	}
	// create a map to store integer
	targets := make(map[int]int, len(nums))
	for k, v := range nums {
		sub := target - v
		// assert map contains sub
		if _, ok := targets[sub]; ok {
			// return index
			return []int{targets[sub], k}
		}
		targets[v] = k
	}
	// can not find the correct result
	return nil
}

package main

import "fmt"

/**
题目：Given an array of characters where each character represents a fruit tree,
you are given two baskets and your goal is to put maximum number of fruits in each basket.
The only restriction is that each basket can have only one type of fruit.

You can start with any tree, but once you have started you can’t skip a tree.
You will pick one fruit from each tree until you cannot, i.e.,
you will stop when you have to pick from a third fruit type.

Write a function to return the maximum number of fruits in both the baskets.

解析：给定一个字符串，每个字符表示一种水果。此时你有两个篮子，并且每个篮子只能装一种类型的水果。
写一个函数来返回两个篮子里的水果的最大数量。这里有几个隐含的条件：
	1：只有两个篮子，意味着最多不能超过2种水果
	2：一旦开始就不能停止，不能跳过，意味着必须是连续的子串（滑动窗口算法的明显标志）
	3：求两个篮子中最多能够装多少个水果？这里是求最大长度

举例：
Example 1:

Input: Fruit=["A", "B", "C", "A", "C"]
Output: 3
Explanation: We can put 2 "C" in one basket and one "A" in the other from the subarray ["C", "A", "C"]
Example 2:

Input: Fruit=["A", "B", "C", "B", "B", "C"]
Output: 5
Explanation: We can put 3 "B" in one basket and two "C" in the other basket.
This can be done if we start with the second letter: ["B", "C", "B", "B", "C"]

模板：
func findLength(strs []string) int
*/

/**
思路1：滑动窗口法
这是 longest substring with k distinct characters 同类型的题目，只不过k没有明说，这里的k=2

time complexity: O(n)
space complexity: O(k)
*/
func findLength(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	var windowStart int
	var charFrequencyMap = make(map[string]int, len(strs))
	var maxLength int

	for windowEnd := 0; windowEnd < len(strs); windowEnd++ {
		charFrequencyMap[strs[windowEnd]]++
		for len(charFrequencyMap) > 2 {
			if _, ok := charFrequencyMap[strs[windowStart]]; ok {
				charFrequencyMap[strs[windowStart]]--
				if charFrequencyMap[strs[windowStart]] == 0 {
					delete(charFrequencyMap, strs[windowStart])
				}
			}
			windowStart++
		}
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}

	return maxLength
}

func main() {
	strs := []string{"A", "B", "C", "A", "C"}
	strs2 := []string{"A", "B", "C", "B", "B", "C"}
	fmt.Println(findLength(strs))
	fmt.Println(findLength(strs2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

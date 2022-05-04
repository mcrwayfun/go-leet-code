package main

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func singleNumber(nums []int) int
*/

/**
解题思路：题目中有一个数字出现了一次，其他数字全部出现两次，假设有数组[x,x,y], 可以使用
2(x+y)-(x+x+y)=y 的方式来求的差异的数字。
1：申明sum和uniqueSum，sum用来计算原数组的和，uniqueSum计算数字出现一次的情况下的和
2：使用一个set来去重，可以uniqueSum

time complexity: O(n)
space complexity: O(n)
*/
func singleNumber(nums []int) int {
	var duplicatedMap = make(map[int]struct{}, len(nums))
	var sum int
	var uniqueSum int
	for _, num := range nums {
		if _, ok := duplicatedMap[num]; !ok {
			uniqueSum += num
			duplicatedMap[num] = struct{}{}
		}
		sum += num
	}
	return uniqueSum*2 - sum
}

/**
解题思路：可以使用异或来解决该问题，它有2个特性非常符合这道题：
1：x^x=0, 即自己异或自己得到的是0
2：0^0^...^x=x, 即0和任何数异或，最终得到是它自己
所以可以得到的结论：
1：成对的数异或后得到的是0
2：0和单身数异或后得到的是单身数

time complexity: O(n)
space complexity: O(1)
*/
func singleNumberWithXOR(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

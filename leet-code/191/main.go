package main

/**
编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。

提示：

请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的示例 3中，输入表示有符号整数 -3。

示例 1：
输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011中，共有三位为 '1'。

示例 2：
输入：00000000000000000000000010000000
输出：1
解释：输入的二进制串 00000000000000000000000010000000中，共有一位为 '1'。

示例 3：
输入：11111111111111111111111111111101
输出：31
解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func hammingWeight(num uint32) int {}
*/

// time complexity: O(m)
// space complexity: O(1)
func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	var countOne int
	for num != 0 {
		if num&1 == 1 {
			countOne++
		}
		num >>= 1
	}
	return countOne
}

/**
解法：有一个特性num & (num-1) 可以去掉低位的1
假设有数字12，二进制为1100，
1: 1100 & 1011 = 1000 (去掉了1100中的低位1）
2：1000 & 0111 = 0000 (去掉了1000中的低位1）
此时结果为0000，结束循环。
 */
// time complexity: O(k), k=二进制1的个数
// space complexity: O(1)
func hammingWeight2(num uint32) int {
	if num == 0 {
		return 0
	}

	var countOne int
	for num != 0 {
		countOne++
		num = num & (num - 1)
	}
	return countOne
}

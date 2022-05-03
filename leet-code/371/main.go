package main

/**
给你两个整数 a 和 b ，不使用 运算符+ 和-，计算并返回两整数之和。

示例 1：
输入：a = 1, b = 2
输出：3

示例 2：
输入：a = 2, b = 3
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-two-integers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func getSum(a int, b int) int
*/

/**
这道题不能使用加法或者减法，所以第一时间要想到位运算。
假设有两个数a=9和b=11，使用二进制的方式表示即为：
	a 1 0 0 1
	b 1 0 1 1
=   1 0 1 0 0

1：假设不考虑进位的情况，仅考虑相加，有4种case
0 0 0
0 1 1
1 0 1
1 1 0
这很明显是a^b(异或)得出的结果

2：假设仅考虑进位的情况，也是有4种case
0 0 0
0 1 0
1 0 0
1 1 1
很明显这是a&b(与)得出的结果

3：假设sum=a^b, carry=a&b<<1(由于要进位，所以有左移1位), 分别来看a^b 和 a&b 得出的结果。
a^b  0010 -> 10000 -> 10100
a&b 10010    00100    00000
由于carry会不断的左移，不断地向carry中填充数字0，所以最终carry会为0, 可以用这个判断是否结束

Time Complexity(m) // m为a和b的二进制长度
Space Complexity: O(1)
*/
func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := a & b << 1

		a = sum
		b = carry
	}
	return a
}

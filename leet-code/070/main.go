package main

/**
假设你正在爬楼梯。需要 n阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func climbStairs(n int) int {}
*/

/**
解题思路1：使用递归的方式来解决，可以得到 fn=f(n-1)+f(n-2)
可以看成这是一颗高度为n的二叉树，而满二叉树的节点为2^n-1.

time complexity: O(2^n)
space complexity: O(2^n)
*/
func climbStairsRecursive(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairsRecursive(n-1) + climbStairsRecursive(n-2)
}

/**
解题思路2：递归过程中产生的变量可以记录下来，可以使用循环的方式来替代。
假设有两个变量first=1, second=1(均表示走上第一个台阶有1种方式）
由于要满足公式f(n)=f(n-1)+f(n-2), 则可以有
third=first+second
first=second
second=third

time complexity: O(n)
space complexity: O(1)
 */
func climbStairsIterator(n int) int {
	var first, second = 1, 1
	for i:=1; i<n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}



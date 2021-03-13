package main

import "fmt"

//一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//示例 1：
//
//输入：n = 2
//输出：2
//示例 2：
//
//输入：n = 7
//输出：21
//示例 3：
//
//输入：n = 0
//输出：1

func main() {
	fmt.Println(numWays(0))
}

func numWays(n int) int {
	if n <= 0 {
		return 1
	}
	if n < 2 {
		return n
	}

	var pre, cur = 1, 2
	for i := 2; i < n; i++ {
		pre, cur = cur%(1e9+7), (pre+cur)%(1e9+7)
	}

	return cur
}
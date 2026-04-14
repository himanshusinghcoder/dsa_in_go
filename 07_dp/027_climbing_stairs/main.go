// Problem: CLIMBING STAIRS
// LeetCode #70 | Difficulty: Easy
// ---------------------------------
// n steps to reach top. Each time climb 1 or 2 steps.
// How many distinct ways to climb to the top?
//
// Example 1:
//   Input:  n = 2
//   Output: 2
//
// Example 2:
//   Input:  n = 3
//   Output: 3
//
// Constraints:
//   1 <= n <= 45
//
// Approach: DP Fibonacci - O(n) time, O(1) space
//   ways(n) = ways(n-1) + ways(n-2)
package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	fmt.Println(climbStairs(1))  // 1
	fmt.Println(climbStairs(2))  // 2
	fmt.Println(climbStairs(3))  // 3
	fmt.Println(climbStairs(5))  // 8
	fmt.Println(climbStairs(10)) // 89
}

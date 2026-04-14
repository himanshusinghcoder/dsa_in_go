// Problem: HOUSE ROBBER
// LeetCode #198 | Difficulty: Medium
// ------------------------------------
// Rob houses (no two adjacent). Return max money.
//
// Example 1:
//   Input:  nums = [1,2,3,1]
//   Output: 4
//
// Example 2:
//   Input:  nums = [2,7,9,3,1]
//   Output: 12
//
// Constraints:
//   1 <= nums.length <= 100
//   0 <= nums[i] <= 400
//
// Approach: DP two variables - O(n) time, O(1) space
//   dp[i] = max(dp[i-1], dp[i-2] + nums[i])
package main

import "fmt"

func rob(nums []int) int {
	prev, curr := 0, 0
	for _, n := range nums {
		prev, curr = curr, max(curr, prev+n)
	}
	return curr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))    // 4
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) // 12
	fmt.Println(rob([]int{5}))              // 5
}

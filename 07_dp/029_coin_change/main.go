// Problem: COIN CHANGE
// LeetCode #322 | Difficulty: Medium
// ------------------------------------
// Minimum coins to make amount. Unlimited use of each coin. Return -1 if impossible.
//
// Example 1:
//   Input:  coins = [1,2,5], amount = 11
//   Output: 3  (5+5+1)
//
// Example 2:
//   Input:  coins = [2], amount = 3
//   Output: -1
//
// Constraints:
//   1 <= coins.length <= 12
//   0 <= amount <= 10^4
//
// Approach: Bottom-up DP - O(amount * coins) time, O(amount) space
//   dp[i] = min coins for amount i. dp[0]=0.
package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))      // 3
	fmt.Println(coinChange([]int{2}, 3))              // -1
	fmt.Println(coinChange([]int{1}, 0))              // 0
	fmt.Println(coinChange([]int{1, 5, 10, 25}, 41)) // 4
}

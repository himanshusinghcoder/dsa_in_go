// Problem: LONGEST COMMON SUBSEQUENCE
// LeetCode #1143 | Difficulty: Medium
// ---------------------------------------
// Return LCS length of text1 and text2.
// A subsequence: delete some chars (maintaining order) without reordering.
//
// Example 1:
//   Input:  text1 = "abcde", text2 = "ace"
//   Output: 3  ("ace")
//
// Example 2:
//   Input:  text1 = "abc", text2 = "abc"
//   Output: 3
//
// Example 3:
//   Input:  text1 = "abc", text2 = "def"
//   Output: 0
//
// Constraints:
//   1 <= text1.length, text2.length <= 1000
//
// Approach: 2D DP - O(m*n) time, O(m*n) space
//   dp[i][j] = LCS of first i chars of text1 and first j of text2.
package main

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))      // 3
	fmt.Println(longestCommonSubsequence("abc", "abc"))        // 3
	fmt.Println(longestCommonSubsequence("abc", "def"))        // 0
	fmt.Println(longestCommonSubsequence("ezupkr", "ubmrapg")) // 2
}

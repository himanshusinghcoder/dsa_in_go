// Problem: LONGEST SUBSTRING WITHOUT REPEATING CHARACTERS
// LeetCode #3 | Difficulty: Medium
// ---------------------------------------------------------
// Find the length of the longest substring without repeating characters.
//
// Example 1:
//   Input:  s = "abcabcbb"
//   Output: 3   ("abc")
//
// Example 2:
//   Input:  s = "bbbbb"
//   Output: 1   ("b")
//
// Example 3:
//   Input:  s = "pwwkew"
//   Output: 3   ("wke")
//
// Constraints:
//   0 <= s.length <= 5*10^4
//   s consists of English letters, digits, symbols, spaces.
//
// Approach: Sliding Window with HashMap - O(n) time, O(min(n,m)) space
//   Expand right; on duplicate, jump left past the previous occurrence.
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	best, l := 0, 0
	for r := 0; r < len(s); r++ {
		if pos, ok := lastSeen[s[r]]; ok && pos >= l {
			l = pos + 1
		}
		lastSeen[s[r]] = r
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring(""))          // 0
}

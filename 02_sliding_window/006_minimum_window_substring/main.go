// Problem: MINIMUM WINDOW SUBSTRING
// LeetCode #76 | Difficulty: Hard
// -----------------------------------
// Find the minimum window in s such that every character in t is included.
// Return "" if no such window exists.
//
// Example 1:
//   Input:  s = "ADOBECODEBANC", t = "ABC"
//   Output: "BANC"
//
// Example 2:
//   Input:  s = "a", t = "a"
//   Output: "a"
//
// Example 3:
//   Input:  s = "a", t = "aa"
//   Output: ""
//
// Constraints:
//   1 <= s.length, t.length <= 10^5
//   s and t consist of uppercase and lowercase English letters.
//
// Approach: Sliding Window - O(m+n) time, O(m+n) space
//   Track 'have' satisfied chars vs 'total' required.
//   Expand right; when fully satisfied shrink left recording best window.
package main

import "fmt"

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	window := make(map[byte]int)
	have, total := 0, len(need)
	best := [2]int{-1, 0}
	l := 0
	for r := 0; r < len(s); r++ {
		c := s[r]
		window[c]++
		if n, ok := need[c]; ok && window[c] == n {
			have++
		}
		for have == total {
			if best[0] == -1 || r-l+1 < best[1]-best[0]+1 {
				best = [2]int{l, r}
			}
			window[s[l]]--
			if n, ok := need[s[l]]; ok && window[s[l]] < n {
				have--
			}
			l++
		}
	}
	if best[0] == -1 {
		return ""
	}
	return s[best[0] : best[1]+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // BANC
	fmt.Println(minWindow("a", "a"))               // a
	fmt.Println(minWindow("a", "aa"))              // (empty)
}

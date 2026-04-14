// Problem: PERMUTATION IN STRING
// LeetCode #567 | Difficulty: Medium
// ------------------------------------
// Return true if s2 contains a permutation of s1.
//
// Example 1:
//   Input:  s1 = "ab", s2 = "eidbaooo"
//   Output: true   ("ba" is in s2)
//
// Example 2:
//   Input:  s1 = "ab", s2 = "eidboaoo"
//   Output: false
//
// Constraints:
//   1 <= s1.length, s2.length <= 10^4
//   s1 and s2 consist of lowercase English letters.
//
// Approach: Fixed-size Window with char frequency - O(n) time, O(1) space
//   Maintain freq arrays for s1 and the current window in s2.
//   Track how many of 26 slots match; if all 26 match -> true.
package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var need, have [26]int
	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
		have[s2[i]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if need[i] == have[i] {
			matches++
		}
	}
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		in := s2[r] - 'a'
		if have[in] == need[in] {
			matches--
		}
		have[in]++
		if have[in] == need[in] {
			matches++
		}
		out := s2[r-len(s1)] - 'a'
		if have[out] == need[out] {
			matches--
		}
		have[out]--
		if have[out] == need[out] {
			matches++
		}
	}
	return matches == 26
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
	fmt.Println(checkInclusion("ab", "eidboaoo")) // false
	fmt.Println(checkInclusion("adc", "dcda"))    // true
}

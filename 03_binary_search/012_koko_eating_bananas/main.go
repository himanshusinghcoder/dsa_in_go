// Problem: KOKO EATING BANANAS
// LeetCode #875 | Difficulty: Medium
// -------------------------------------
// n piles of bananas. Guards return in h hours. Koko eats k bananas/hour.
// Find the minimum integer k to eat all bananas within h hours.
//
// Example 1:
//   Input:  piles = [3,6,7,11], h = 8
//   Output: 4
//
// Example 2:
//   Input:  piles = [30,11,23,4,20], h = 5
//   Output: 30
//
// Example 3:
//   Input:  piles = [30,11,23,4,20], h = 6
//   Output: 23
//
// Constraints:
//   1 <= piles.length <= h <= 10^9
//
// Approach: Binary Search on answer - O(n log m) where m = max pile
//   Search k in [1, max(piles)]. For each k compute total hours (sum of ceil(p/k)).
package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	for _, p := range piles {
		if p > hi {
			hi = p
		}
	}
	for lo < hi {
		mid := lo + (hi-lo)/2
		hours := 0
		for _, p := range piles {
			hours += (p + mid - 1) / mid
		}
		if hours <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))      // 4
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)) // 30
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)) // 23
}

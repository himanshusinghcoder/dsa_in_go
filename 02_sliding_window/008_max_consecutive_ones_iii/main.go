// Problem: MAX CONSECUTIVE ONES III
// LeetCode #1004 | Difficulty: Medium
// ---------------------------------------
// Binary array nums; flip at most k zeros. Return max consecutive 1s.
//
// Example 1:
//   Input:  nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
//   Output: 6
//
// Example 2:
//   Input:  nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
//   Output: 10
//
// Constraints:
//   1 <= nums.length <= 10^5
//   nums[i] is 0 or 1.
//   0 <= k <= nums.length
//
// Approach: Sliding Window - O(n) time, O(1) space
//   Expand right counting zeros. When zeros > k, shrink from left.
package main

import "fmt"

func longestOnes(nums []int, k int) int {
	l, zeros, best := 0, 0, 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[l] == 0 {
				zeros--
			}
			l++
		}
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))                           // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
}

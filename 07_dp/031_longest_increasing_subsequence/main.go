// Problem: LONGEST INCREASING SUBSEQUENCE
// LeetCode #300 | Difficulty: Medium
// -----------------------------------------
// Return length of the longest strictly increasing subsequence.
//
// Example 1:
//   Input:  nums = [10,9,2,5,3,7,101,18]
//   Output: 4  ([2,3,7,101])
//
// Example 2:
//   Input:  nums = [0,1,0,3,2,3]
//   Output: 4
//
// Constraints:
//   1 <= nums.length <= 2500
//   -10^4 <= nums[i] <= 10^4
//
// Approach: Patience sort (binary search) - O(n log n) time, O(n) space
//   tails[i] = smallest tail for increasing subseq of length i+1.
//   Binary search to extend or replace in tails.
package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, n := range nums {
		pos := sort.SearchInts(tails, n)
		if pos == len(tails) {
			tails = append(tails, n)
		} else {
			tails[pos] = n
		}
	}
	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))            // 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7}))               // 1
}

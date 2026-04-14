// Problem: FIND MINIMUM IN ROTATED SORTED ARRAY
// LeetCode #153 | Difficulty: Medium
// -----------------------------------------------
// Sorted array rotated 1-n times. Find the minimum element. O(log n).
//
// Example 1:
//   Input:  nums = [3,4,5,1,2]
//   Output: 1
//
// Example 2:
//   Input:  nums = [4,5,6,7,0,1,2]
//   Output: 0
//
// Example 3:
//   Input:  nums = [11,13,15,17]
//   Output: 11
//
// Constraints:
//   1 <= n <= 5000, all unique.
//
// Approach: Binary Search - O(log n) time
//   If nums[mid] > nums[hi], minimum is in right half. Else go left.
package main

import "fmt"

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return nums[lo]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
	fmt.Println(findMin([]int{11, 13, 15, 17}))       // 11
}

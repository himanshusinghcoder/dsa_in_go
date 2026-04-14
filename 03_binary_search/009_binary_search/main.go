// Problem: BINARY SEARCH
// LeetCode #704 | Difficulty: Easy
// ----------------------------------
// Given a sorted array nums and target, return the index of target or -1.
// Must run in O(log n).
//
// Example 1:
//   Input:  nums = [-1,0,3,5,9,12], target = 9
//   Output: 4
//
// Example 2:
//   Input:  nums = [-1,0,3,5,9,12], target = 2
//   Output: -1
//
// Constraints:
//   1 <= nums.length <= 10^4
//   All integers unique. Sorted ascending.
//
// Approach: Classic Binary Search - O(log n) time, O(1) space
package main

import "fmt"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
	fmt.Println(search([]int{5}, 5))                    // 0
}

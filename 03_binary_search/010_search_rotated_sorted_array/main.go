// Problem: SEARCH IN ROTATED SORTED ARRAY
// LeetCode #33 | Difficulty: Medium
// -----------------------------------------
// Sorted array possibly rotated at some pivot. Find target or return -1.
// Must be O(log n).
//
// Example 1:
//   Input:  nums = [4,5,6,7,0,1,2], target = 0
//   Output: 4
//
// Example 2:
//   Input:  nums = [4,5,6,7,0,1,2], target = 3
//   Output: -1
//
// Constraints:
//   1 <= nums.length <= 5000
//   All values are unique.
//
// Approach: Binary Search with rotation detection - O(log n) time
//   Determine which half is sorted. If target falls in sorted half, go there.
package main

import "fmt"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half sorted
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half sorted
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1
	fmt.Println(search([]int{3, 1}, 1))                  // 1
}

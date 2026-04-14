// Problem: THREE SUM
// LeetCode #15 | Difficulty: Medium
// ------------------------------------
// Given nums, return all triplets [i,j,k] such that nums[i]+nums[j]+nums[k]==0.
// The solution set must not contain duplicate triplets.
//
// Example 1:
//   Input:  nums = [-1,0,1,2,-1,-4]
//   Output: [[-1,-1,2],[-1,0,1]]
//
// Example 2:
//   Input:  nums = [0,0,0]
//   Output: [[0,0,0]]
//
// Constraints:
//   3 <= nums.length <= 3000
//   -10^5 <= nums[i] <= 10^5
//
// Approach: Sort + Two Pointers - O(n^2) time, O(1) extra space
//   Sort. Fix nums[i], use l/r pointers for sum == -nums[i].
//   Skip duplicates at both levels.
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s == 0:
				result = append(result, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case s < 0:
				l++
			default:
				r--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{0, 1, 1}))              // []
	fmt.Println(threeSum([]int{0, 0, 0}))              // [[0 0 0]]
}

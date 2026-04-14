// Problem: TWO SUM
// LeetCode #1 | Difficulty: Easy
// --------------------------------
// Given an array of integers nums and an integer target, return indices of
// the two numbers such that they add up to target. Each input has exactly
// one solution. You may not use the same element twice.
// Return the answer in any order.
//
// Example 1:
//   Input:  nums = [2,7,11,15], target = 9
//   Output: [0,1]
//   Explanation: nums[0] + nums[1] == 9
//
// Example 2:
//   Input:  nums = [3,2,4], target = 6
//   Output: [1,2]
//
// Example 3:
//   Input:  nums = [3,3], target = 6
//   Output: [0,1]
//
// Constraints:
//   2 <= nums.length <= 10^4
//   -10^9 <= nums[i] <= 10^9
//   -10^9 <= target <= 10^9
//   Only one valid answer exists.
//
// Approach: Hash Map - O(n) time, O(n) space
//   For each number, look up (target - num) in the map.
//   If found return both indices. Otherwise store num -> index.
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))       // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))           // [0 1]
}

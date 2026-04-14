// Problem: SUBSETS
// LeetCode #78 | Difficulty: Medium
// -----------------------------------
// Return all possible subsets (power set) of unique nums.
// No duplicate subsets.
//
// Example 1:
//   Input:  nums = [1,2,3]
//   Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
// Example 2:
//   Input:  nums = [0]
//   Output: [[],[0]]
//
// Constraints:
//   1 <= nums.length <= 10, all unique.
//
// Approach: Backtracking - O(2^n * n) time
//   Record path at every call. Iterate from start to avoid reuse.
package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var bt func(start int, path []int)
	bt = func(start int, path []int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			bt(i+1, path)
			path = path[:len(path)-1]
		}
	}
	bt(0, []int{})
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	// [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsets([]int{0}))
	// [[] [0]]
}

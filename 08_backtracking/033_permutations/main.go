// Problem: PERMUTATIONS
// LeetCode #46 | Difficulty: Medium
// -----------------------------------
// Return all permutations of distinct integers.
//
// Example 1:
//   Input:  nums = [1,2,3]
//   Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// Example 2:
//   Input:  nums = [0,1]
//   Output: [[0,1],[1,0]]
//
// Constraints:
//   1 <= nums.length <= 6, all unique.
//
// Approach: Backtracking with used[] - O(n! * n) time
package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	var bt func(path []int)
	bt = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i, n := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, n)
			bt(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	bt([]int{})
	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3})) // 6 permutations
	fmt.Println(permute([]int{0, 1}))    // [[0 1] [1 0]]
}

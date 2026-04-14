// Problem: COMBINATION SUM
// LeetCode #39 | Difficulty: Medium
// -----------------------------------
// Distinct candidates; find all combinations that sum to target.
// Same number may be used unlimited times.
//
// Example 1:
//   Input:  candidates = [2,3,6,7], target = 7
//   Output: [[2,2,3],[7]]
//
// Example 2:
//   Input:  candidates = [2,3], target = 6
//   Output: [[2,2,2],[3,3]]
//
// Constraints:
//   1 <= candidates.length <= 30, 2 <= candidates[i] <= 40, 1 <= target <= 40.
//
// Approach: Backtracking - O(N^(T/M))
//   Allow reuse: recurse with same index i. Prune if candidate > remaining.
package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var bt func(start, rem int, path []int)
	bt = func(start, rem int, path []int) {
		if rem == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > rem {
				continue
			}
			path = append(path, candidates[i])
			bt(i, rem-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	bt(0, target, []int{})
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]
	fmt.Println(combinationSum([]int{2, 3}, 6))        // [[2 2 2] [3 3]]
	fmt.Println(combinationSum([]int{2}, 1))            // []
}

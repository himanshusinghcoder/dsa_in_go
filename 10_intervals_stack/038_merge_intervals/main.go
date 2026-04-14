// Problem: MERGE INTERVALS
// LeetCode #56 | Difficulty: Medium
// -----------------------------------
// Merge all overlapping intervals.
//
// Example 1:
//   Input:  intervals = [[1,3],[2,6],[8,10],[15,18]]
//   Output: [[1,6],[8,10],[15,18]]
//
// Example 2:
//   Input:  intervals = [[1,4],[4,5]]
//   Output: [[1,5]]
//
// Constraints:
//   1 <= intervals.length <= 10^4
//   0 <= starti <= endi <= 10^4
//
// Approach: Sort by start + linear scan - O(n log n) time
//   If current overlaps last merged (start <= last.end), extend last.end.
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for _, cur := range intervals[1:] {
		last := result[len(result)-1]
		if cur[0] <= last[1] {
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			result = append(result, cur)
		}
	}
	return result
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                     // [[1 5]]
}

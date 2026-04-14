// Problem: TRAPPING RAIN WATER
// LeetCode #42 | Difficulty: Hard
// ---------------------------------
// Given n non-negative integers representing an elevation map (width=1 per bar),
// compute how much water it can trap after raining.
//
// Example 1:
//   Input:  height = [0,1,0,2,1,0,1,3,2,1,2,1]
//   Output: 6
//
// Example 2:
//   Input:  height = [4,2,0,3,2,5]
//   Output: 9
//
// Constraints:
//   n == height.length, 1 <= n <= 2*10^4
//   0 <= height[i] <= 10^5
//
// Approach: Two Pointers - O(n) time, O(1) space
//   water[i] = min(maxLeft, maxRight) - height[i]
//   Process the side with the smaller max: guaranteed water is bounded by that max.
package main

import "fmt"

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax, water := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				water += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				water += rightMax - height[r]
			}
			r--
		}
	}
	return water
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                     // 9
}

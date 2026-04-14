// Problem: CONTAINER WITH MOST WATER
// LeetCode #11 | Difficulty: Medium
// ------------------------------------
// n vertical lines at positions i with heights height[i].
// Find two lines forming a container that holds the most water.
// Return the maximum volume.
//
// Example 1:
//   Input:  height = [1,8,6,2,5,4,8,3,7]
//   Output: 49
//
// Example 2:
//   Input:  height = [1,1]
//   Output: 1
//
// Constraints:
//   n == height.length, 2 <= n <= 10^5
//   0 <= height[i] <= 10^4
//
// Approach: Two Pointers - O(n) time, O(1) space
//   Start widest (l=0, r=n-1). Move whichever side is shorter inward.
package main

import "fmt"

func maxArea(height []int) int {
	l, r, best := 0, len(height)-1, 0
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		if area := h * (r - l); area > best {
			best = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                        // 1
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))              // 16
}

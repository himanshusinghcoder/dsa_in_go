// Problem: LARGEST RECTANGLE IN HISTOGRAM
// LeetCode #84 | Difficulty: Hard
// -----------------------------------------
// Return the area of the largest rectangle in the histogram.
//
// Example 1:
//   Input:  heights = [2,1,5,6,2,3]
//   Output: 10
//
// Example 2:
//   Input:  heights = [2,4]
//   Output: 4
//
// Constraints:
//   1 <= heights.length <= 10^5
//   0 <= heights[i] <= 10^4
//
// Approach: Monotonic Stack (increasing) - O(n) time, O(n) space
//   When we see a shorter bar, pop taller bars and compute their max rectangle.
//   width = i - previous_stack_index - 1  (or i if stack empty).
package main

import "fmt"

func largestRectangleArea(heights []int) int {
	stack := []int{}
	best := 0
	n := len(heights)
	for i := 0; i <= n; i++ {
		h := 0
		if i < n {
			h = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}
			if heights[top]*w > best {
				best = heights[top] * w
			}
		}
		stack = append(stack, i)
	}
	return best
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))        // 10
	fmt.Println(largestRectangleArea([]int{2, 4}))                     // 4
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3})) // 16
}

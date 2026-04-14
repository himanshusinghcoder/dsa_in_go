// Problem: DAILY TEMPERATURES
// LeetCode #739 | Difficulty: Medium
// ------------------------------------
// Return answer[i] = days to wait for a warmer temperature after day i.
// 0 if no such day.
//
// Example 1:
//   Input:  temperatures = [73,74,75,71,69,72,76,73]
//   Output: [1,1,4,2,1,1,0,0]
//
// Example 2:
//   Input:  temperatures = [30,40,50,60]
//   Output: [1,1,1,0]
//
// Constraints:
//   1 <= temperatures.length <= 10^5
//   30 <= temperatures[i] <= 100
//
// Approach: Monotonic Stack (decreasing) - O(n) time, O(n) space
//   Stack of indices. For each temp, pop indices with cooler temp; record gap.
package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // [1 1 4 2 1 1 0 0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                  // [1 1 1 0]
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))                      // [1 1 0]
}

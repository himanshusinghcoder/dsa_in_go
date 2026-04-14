// Problem: COURSE SCHEDULE
// LeetCode #207 | Difficulty: Medium
// ------------------------------------
// Can you finish all numCourses given prerequisites?
// prerequisites[i] = [a, b] means b must be taken before a.
//
// Example 1:
//   Input:  numCourses = 2, prerequisites = [[1,0]]
//   Output: true
//
// Example 2:
//   Input:  numCourses = 2, prerequisites = [[1,0],[0,1]]
//   Output: false (cycle)
//
// Constraints:
//   1 <= numCourses <= 2000
//   0 <= prerequisites.length <= 5000
//
// Approach: DFS cycle detection - O(V+E) time
//   State 0=unvisited, 1=visiting, 2=done. Back edge (state=1) = cycle.
package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
	}
	state := make([]int, numCourses)
	var dfs func(int) bool
	dfs = func(node int) bool {
		if state[node] == 1 {
			return true
		}
		if state[node] == 2 {
			return false
		}
		state[node] = 1
		for _, nb := range adj[node] {
			if dfs(nb) {
				return true
			}
		}
		state[node] = 2
		return false
	}
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 && dfs(i) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))          // true
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))  // false
	fmt.Println(canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}})) // true
}

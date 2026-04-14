// Problem: ROTTING ORANGES
// LeetCode #994 | Difficulty: Medium
// ------------------------------------
// Grid: 0=empty, 1=fresh, 2=rotten. Each minute rotten spreads to adjacent fresh.
// Return min minutes until all oranges rot, or -1 if impossible.
//
// Example 1:
//   Input:  grid = [[2,1,1],[1,1,0],[0,1,1]]
//   Output: 4
//
// Example 2:
//   Input:  grid = [[2,1,1],[0,1,1],[1,0,1]]
//   Output: -1
//
// Example 3:
//   Input:  grid = [[0,2]]
//   Output: 0
//
// Constraints:
//   1 <= m, n <= 10; grid[i][j] in {0,1,2}.
//
// Approach: Multi-source BFS - O(m*n) time, O(m*n) space
package main

import "fmt"

type pt struct{ r, c int }

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	var q []pt
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				fresh++
			} else if grid[r][c] == 2 {
				q = append(q, pt{r, c})
			}
		}
	}
	dirs := []pt{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(q) > 0 && fresh > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			p := q[i]
			for _, d := range dirs {
				nr, nc := p.r+d.r, p.c+d.c
				if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					q = append(q, pt{nr, nc})
				}
			}
		}
		q = q[size:]
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})) // -1
	fmt.Println(orangesRotting([][]int{{0, 2}}))                            // 0
}

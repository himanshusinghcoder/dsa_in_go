// Problem: NUMBER OF ISLANDS
// LeetCode #200 | Difficulty: Medium
// ------------------------------------
// Count islands in a grid of '1' (land) and '0' (water).
// Islands are formed by connecting land cells horizontally/vertically.
//
// Example 1:
//   Input:  grid = [["1","1","1","1","0"],["1","1","0","1","0"],
//                   ["1","1","0","0","0"],["0","0","0","0","0"]]
//   Output: 1
//
// Example 2:
//   Input:  grid = [["1","1","0","0","0"],["1","1","0","0","0"],
//                   ["0","0","1","0","0"],["0","0","0","1","1"]]
//   Output: 3
//
// Constraints:
//   1 <= m, n <= 300; grid[i][j] is '0' or '1'.
//
// Approach: DFS flood-fill - O(m*n) time, O(m*n) space
//   Find a '1', increment count, DFS to mark entire island as '0'.
package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	count := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}
	return count
}

func main() {
	g1 := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(g1)) // 1

	g2 := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(g2)) // 3
}

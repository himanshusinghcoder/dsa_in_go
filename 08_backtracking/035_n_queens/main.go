// Problem: N-QUEENS
// LeetCode #51 | Difficulty: Hard
// ---------------------------------
// Place n queens on n x n board so no two attack each other.
// Return all distinct board configurations.
//
// Example 1:
//   Input:  n = 4
//   Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//
// Example 2:
//   Input:  n = 1
//   Output: [["Q"]]
//
// Constraints:
//   1 <= n <= 9
//
// Approach: Backtracking row-by-row - O(n!) time
//   Track occupied cols, diagonals (row-col), anti-diagonals (row+col).
package main

import "fmt"

func solveNQueens(n int) [][]string {
	var result [][]string
	cols := make(map[int]bool)
	diag := make(map[int]bool)
	adiag := make(map[int]bool)
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	var bt func(row int)
	bt = func(row int) {
		if row == n {
			snap := make([]string, n)
			for i, r := range board {
				snap[i] = string(r)
			}
			result = append(result, snap)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag[row-col] || adiag[row+col] {
				continue
			}
			cols[col] = true
			diag[row-col] = true
			adiag[row+col] = true
			board[row][col] = 'Q'
			bt(row + 1)
			board[row][col] = '.'
			cols[col] = false
			diag[row-col] = false
			adiag[row+col] = false
		}
	}
	bt(0)
	return result
}

func main() {
	s4 := solveNQueens(4)
	fmt.Printf("n=4: %d solutions\n", len(s4)) // 2
	for _, sol := range s4 {
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println("---")
	}
	fmt.Printf("n=1: %d solution\n", len(solveNQueens(1))) // 1
}

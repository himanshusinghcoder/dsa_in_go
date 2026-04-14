// Problem: BINARY TREE LEVEL ORDER TRAVERSAL
// LeetCode #102 | Difficulty: Medium
// -------------------------------------------
// Return nodes' values grouped level by level (left to right).
//
// Example 1:
//   Input:  root = [3,9,20,null,null,15,7]
//   Output: [[3],[9,20],[15,7]]
//
// Example 2:
//   Input:  root = [1]
//   Output: [[1]]
//
// Constraints:
//   0 <= number of nodes <= 2000
//   -1000 <= Node.val <= 1000
//
// Approach: BFS with level snapshot - O(n) time, O(n) space
//   Snapshot queue size = current level width. Process exactly that many nodes.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := q[i]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
		result = append(result, level)
	}
	return result
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(root))              // [[3] [9 20] [15 7]]
	fmt.Println(levelOrder(&TreeNode{Val: 1})) // [[1]]
	fmt.Println(levelOrder(nil))               // []
}

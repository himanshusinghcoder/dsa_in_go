// Problem: VALIDATE BINARY SEARCH TREE
// LeetCode #98 | Difficulty: Medium
// --------------------------------------
// Determine if the binary tree is a valid BST.
// Valid BST: left < root < right at every node (not just immediate children).
//
// Example 1:
//   Input:  root = [2,1,3]
//   Output: true
//
// Example 2:
//   Input:  root = [5,1,4,null,null,3,6]
//   Output: false
//
// Constraints:
//   1 <= number of nodes <= 10^4
//   -2^31 <= Node.val <= 2^31-1
//
// Approach: DFS with bounds - O(n) time, O(h) space
//   Pass (lo, hi) range to each node. Left child: hi=parent.Val. Right: lo=parent.Val.
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(*TreeNode, int, int) bool
	check = func(node *TreeNode, lo, hi int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lo || node.Val >= hi {
			return false
		}
		return check(node.Left, lo, node.Val) && check(node.Right, node.Val, hi)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}

func main() {
	root1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root1)) // true

	root2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root2)) // false
}

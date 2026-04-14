// Problem: MAXIMUM DEPTH OF BINARY TREE
// LeetCode #104 | Difficulty: Easy
// ----------------------------------------
// Return the maximum depth of a binary tree (longest root-to-leaf path).
//
// Example 1:
//   Input:  root = [3,9,20,null,null,15,7]
//   Output: 3
//
// Example 2:
//   Input:  root = [1,null,2]
//   Output: 2
//
// Constraints:
//   0 <= number of nodes <= 10^4
//   -100 <= Node.val <= 100
//
// Approach: DFS post-order - O(n) time, O(h) space
//   depth(node) = 1 + max(depth(left), depth(right))
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := maxDepth(root.Left), maxDepth(root.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}

func main() {
	root := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20,
			Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(maxDepth(root))                                         // 3
	fmt.Println(maxDepth(&TreeNode{Val: 1, Right: &TreeNode{Val: 2}})) // 2
	fmt.Println(maxDepth(nil))                                          // 0
}

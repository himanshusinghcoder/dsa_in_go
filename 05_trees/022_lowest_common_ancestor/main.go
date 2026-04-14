// Problem: LOWEST COMMON ANCESTOR OF A BINARY TREE
// LeetCode #236 | Difficulty: Medium
// --------------------------------------------------
// Find the LCA of nodes p and q. The LCA is the lowest node that has both
// p and q as descendants (a node can be a descendant of itself).
//
// Example 1:
//   Input:  root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//   Output: 3
//
// Example 2:
//   Input:  root = [3,5,1,...], p = 5, q = 4
//   Output: 5
//
// Constraints:
//   2 <= number of nodes <= 10^5, all values unique.
//
// Approach: DFS post-order - O(n) time, O(h) space
//   If root is nil/p/q return root. Recurse left and right.
//   Both non-nil -> current node is LCA. Otherwise return non-nil side.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {
	root := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 1}
	n4 := &TreeNode{Val: 4}
	root.Left = n5
	root.Right = n1
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = n4
	fmt.Println(lowestCommonAncestor(root, n5, n1).Val) // 3
	fmt.Println(lowestCommonAncestor(root, n5, n4).Val) // 5
}

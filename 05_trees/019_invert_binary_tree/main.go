// Problem: INVERT BINARY TREE
// LeetCode #226 | Difficulty: Easy
// ----------------------------------
// Invert a binary tree (mirror it) and return its root.
//
// Example 1:
//   Input:  root = [4,2,7,1,3,6,9]
//   Output: [4,7,2,9,6,3,1]
//
// Example 2:
//   Input:  root = [2,1,3]
//   Output: [2,3,1]
//
// Constraints:
//   0 <= number of nodes <= 100
//   -100 <= Node.val <= 100
//
// Approach: DFS - O(n) time, O(h) space
//   Swap left and right children, then recurse on each.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func collect(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7,
			Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	fmt.Println(collect(invertTree(root))) // [4 7 2 9 6 3 1]
}

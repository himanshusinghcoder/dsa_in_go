// Problem: DIAMETER OF BINARY TREE
// LeetCode #543 | Difficulty: Easy
// ----------------------------------
// Return the length of the diameter (longest path between any two nodes).
// The path may not pass through the root. Length = number of edges.
//
// Example 1:
//   Input:  root = [1,2,3,4,5]
//   Output: 3   (path 4-2-1-3 or 5-2-1-3)
//
// Example 2:
//   Input:  root = [1,2]
//   Output: 1
//
// Constraints:
//   1 <= number of nodes <= 10^4
//   -100 <= Node.val <= 100
//
// Approach: DFS with global max - O(n) time, O(h) space
//   diameter through node = leftHeight + rightHeight; track global max.
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var h func(*TreeNode) int
	h = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := h(node.Left), h(node.Right)
		if l+r > best {
			best = l + r
		}
		if l > r {
			return 1 + l
		}
		return 1 + r
	}
	h(root)
	return best
}

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3}}
	fmt.Println(diameterOfBinaryTree(root))                                        // 3
	fmt.Println(diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}})) // 1
}

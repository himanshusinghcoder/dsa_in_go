// Problem: LINKED LIST CYCLE
// LeetCode #141 | Difficulty: Easy
// ----------------------------------
// Determine if the linked list has a cycle. Return true if yes.
//
// Example 1:
//   Input:  head = [3,2,0,-4], pos = 1
//   Output: true
//
// Example 2:
//   Input:  head = [1], pos = -1
//   Output: false
//
// Constraints:
//   0 <= number of nodes <= 10^4
//   -10^5 <= Node.val <= 10^5
//
// Approach: Floyd's Cycle Detection - O(n) time, O(1) space
//   slow moves 1 step, fast moves 2 steps. Meet = cycle.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2 // cycle
	fmt.Println(hasCycle(n1)) // true

	a := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	fmt.Println(hasCycle(a)) // false
}

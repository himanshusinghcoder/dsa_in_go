// Problem: REMOVE NTH NODE FROM END OF LIST
// LeetCode #19 | Difficulty: Medium
// -------------------------------------------
// Remove the n-th node from the end of the linked list and return its head.
//
// Example 1:
//   Input:  head = [1,2,3,4,5], n = 2
//   Output: [1,2,3,5]
//
// Example 2:
//   Input:  head = [1], n = 1
//   Output: []
//
// Constraints:
//   1 <= sz <= 30, 1 <= n <= sz
//
// Approach: Two Pointers one-pass - O(n) time, O(1) space
//   Advance fast n+1 steps ahead (from dummy). Then move both until fast=nil.
//   slow.Next is the target; unlink it.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func build(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {
	fmt.Println(toSlice(removeNthFromEnd(build([]int{1, 2, 3, 4, 5}), 2))) // [1 2 3 5]
	fmt.Println(toSlice(removeNthFromEnd(build([]int{1}), 1)))              // []
	fmt.Println(toSlice(removeNthFromEnd(build([]int{1, 2}), 1)))           // [1]
}

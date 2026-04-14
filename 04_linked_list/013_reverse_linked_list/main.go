// Problem: REVERSE LINKED LIST
// LeetCode #206 | Difficulty: Easy
// ----------------------------------
// Given the head of a singly linked list, reverse it and return the new head.
//
// Example 1:
//   Input:  head = [1,2,3,4,5]
//   Output: [5,4,3,2,1]
//
// Example 2:
//   Input:  head = [1,2]
//   Output: [2,1]
//
// Constraints:
//   0 <= number of nodes <= 5000
//   -5000 <= Node.val <= 5000
//
// Approach: Iterative three-pointer - O(n) time, O(1) space
//   prev=nil, curr=head. Re-link curr.Next=prev, advance both.
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

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	fmt.Println(toSlice(reverseList(build([]int{1, 2, 3, 4, 5})))) // [5 4 3 2 1]
	fmt.Println(toSlice(reverseList(build([]int{1, 2}))))           // [2 1]
	fmt.Println(toSlice(reverseList(nil)))                          // []
}

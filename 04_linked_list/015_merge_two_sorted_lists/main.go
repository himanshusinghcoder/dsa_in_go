// Problem: MERGE TWO SORTED LISTS
// LeetCode #21 | Difficulty: Easy
// ---------------------------------
// Merge two sorted linked lists and return the merged sorted list.
//
// Example 1:
//   Input:  list1 = [1,2,4], list2 = [1,3,4]
//   Output: [1,1,2,3,4,4]
//
// Example 2:
//   Input:  list1 = [], list2 = []
//   Output: []
//
// Constraints:
//   0 <= number of nodes in each list <= 50
//   -100 <= Node.val <= 100
//   Both lists are sorted non-decreasingly.
//
// Approach: Dummy head + compare - O(m+n) time, O(1) space
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func main() {
	fmt.Println(toSlice(mergeTwoLists(build([]int{1, 2, 4}), build([]int{1, 3, 4})))) // [1 1 2 3 4 4]
	fmt.Println(toSlice(mergeTwoLists(nil, nil)))                                       // []
	fmt.Println(toSlice(mergeTwoLists(nil, build([]int{0}))))                           // [0]
}

// Problem: MERGE K SORTED LISTS
// LeetCode #23 | Difficulty: Hard
// ----------------------------------
// Merge k sorted linked lists into one sorted list.
//
// Example:
//   Input:  lists = [[1,4,5],[1,3,4],[2,6]]
//   Output: [1,1,2,3,4,4,5,6]
//
// Constraints:
//   k == lists.length, 0 <= k <= 10^4
//   0 <= lists[i].length <= 500
//
// Approach: Divide and Conquer - O(n log k) time, O(log k) stack
//   Split list of lists in halves, recurse, then merge the two results.
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

func mergeTwo(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	return mergeTwo(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func main() {
	lists := []*ListNode{build([]int{1, 4, 5}), build([]int{1, 3, 4}), build([]int{2, 6})}
	fmt.Println(toSlice(mergeKLists(lists))) // [1 1 2 3 4 4 5 6]
	fmt.Println(toSlice(mergeKLists(nil)))   // []
}

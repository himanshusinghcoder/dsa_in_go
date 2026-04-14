// Problem: KTH LARGEST ELEMENT IN AN ARRAY
// LeetCode #215 | Difficulty: Medium
// ------------------------------------------
// Return the k-th largest element in the array (not k-th distinct).
//
// Example 1:
//   Input:  nums = [3,2,1,5,6,4], k = 2
//   Output: 5
//
// Example 2:
//   Input:  nums = [3,2,3,1,2,4,5,5,6], k = 4
//   Output: 4
//
// Constraints:
//   1 <= k <= nums.length <= 10^5
//   -10^4 <= nums[i] <= 10^4
//
// Approach: Min-heap of size k - O(n log k) time, O(k) space
//   Maintain min-heap of top-k elements. Pop when size > k.
//   Top = k-th largest.
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
}

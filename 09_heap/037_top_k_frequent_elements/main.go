// Problem: TOP K FREQUENT ELEMENTS
// LeetCode #347 | Difficulty: Medium
// -------------------------------------
// Return the k most frequent elements. Must be better than O(n log n).
//
// Example 1:
//   Input:  nums = [1,1,1,2,2,3], k = 2
//   Output: [1,2]
//
// Example 2:
//   Input:  nums = [1], k = 1
//   Output: [1]
//
// Constraints:
//   1 <= nums.length <= 10^5
//   Answer is unique.
//
// Approach: Min-heap by frequency of size k - O(n log k) time
//   Count frequencies. Min-heap on freq; pop when size > k. Remaining = top-k.
package main

import (
	"container/heap"
	"fmt"
)

type entry struct{ val, freq int }
type FreqHeap []entry

func (h FreqHeap) Len() int            { return len(h) }
func (h FreqHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h FreqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *FreqHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *FreqHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	h := &FreqHeap{}
	heap.Init(h)
	for val, f := range freq {
		heap.Push(h, entry{val, f})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(entry).val
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))      // [1 2]
	fmt.Println(topKFrequent([]int{1}, 1))                       // [1]
	fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)) // [-1 2]
}

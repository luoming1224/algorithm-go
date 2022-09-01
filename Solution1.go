package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solution1(A []int) int {
	nh := &IntHeap{}
	heap.Init(nh)

	numLen := len(A)
	sum := 0
	res := 0
	for i := 0; i < numLen; i++ {
		if A[i] < 0 {
			heap.Push(nh, A[i])
		}
		sum += A[i]
		if sum < 0 {
			for sum < 0 {
				if nh.Len() > 0 {
					if v, ok := heap.Pop(nh).(int); ok {
						sum += (-1) * v
						res++
					}
				} else {
					break
				}
			}
		}
	}
	return res
}

func main() {
	A := []int{10, -10, -1, -1, 10} //1

	fmt.Println(solution1(A))

	A = []int{-1, -1, -1, 1, 1, 1, 1} //3
	fmt.Println(solution1(A))

	A = []int{5, -2, -3, 1} //0
	fmt.Println(solution1(A))
}

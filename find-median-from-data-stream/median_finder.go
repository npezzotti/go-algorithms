package median_finder

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() any {
	return (*h)[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() any {
	return (*h)[0]
}

type MedianFinder struct {
	higherHalf *MinHeap
	lowerHalf  *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		higherHalf: &MinHeap{},
		lowerHalf:  &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.lowerHalf, num)

	if this.lowerHalf.Len() > 0 && this.higherHalf.Len() > 0 &&
		this.lowerHalf.Top().(int) > this.higherHalf.Top().(int) {
		val := heap.Pop(this.lowerHalf)
		heap.Push(this.higherHalf, val)
	}

	if this.lowerHalf.Len() > this.higherHalf.Len()+1 {
		val := heap.Pop(this.lowerHalf)
		heap.Push(this.higherHalf, val)
	}

	if this.higherHalf.Len() > this.lowerHalf.Len()+1 {
		val := heap.Pop(this.higherHalf)
		heap.Push(this.lowerHalf, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lowerHalf.Len() == this.higherHalf.Len() {
		return float64(this.lowerHalf.Top().(int)+this.higherHalf.Top().(int)) / 2.0
	} else if this.lowerHalf.Len() > this.higherHalf.Len() {
		return float64(this.lowerHalf.Top().(int))
	} else {
		return float64(this.higherHalf.Top().(int))
	}
}

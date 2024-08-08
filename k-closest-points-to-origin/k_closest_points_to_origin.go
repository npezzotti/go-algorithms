package k_closest_points_to_origin

import "container/heap"

type Point struct {
	distance int
	x        int
	y        int
}

type PointHeap []Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := new(PointHeap)
	heap.Init(h)

	for _, point := range points {
		d := point[0]*point[0] + point[1]*point[1]
		heap.Push(h, Point{
			distance: d,
			x:        point[0],
			y:        point[1],
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		result[i] = []int{p.x, p.y}
	}

	return result
}

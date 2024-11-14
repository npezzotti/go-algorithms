package task_scheduler

import (
	"container/heap"
)

type Task struct {
	name  string
	count int
}

type PriorityQueue []Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Task))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[0 : n-1]
	return task
}

func leastInterval(tasks []byte, n int) int {
	taskCounts := make([]int, 26)

	for _, task := range tasks {
		taskCounts[task-'A']++
	}

	h := make(PriorityQueue, 0, len(tasks))
	heap.Init(&h)

	for i, freq := range taskCounts {
		if freq > 0 {
			char := byte('A' + i)
			heap.Push(&h, Task{name: string(char), count: freq})
		}
	}

	time := 0
	for h.Len() > 0 {
		var remaining []Task
		cycle := n + 1

		for cycle > 0 && h.Len() > 0 {
			if h.Len() > 0 {
				t := heap.Pop(&h).(Task)
				if t.count > 1 {
					remaining = append(remaining, Task{name: t.name, count: t.count - 1})
				}

				time++
				cycle--
			}
		}

		for _, task := range remaining {
			heap.Push(&h, task)
		}

		if h.Len() > 0 {
			time += cycle
		}
	}

	return time
}

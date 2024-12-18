package top_k_frequent_words

import (
	"container/heap"
)

type word struct {
	word  string
	count int
}

type WordHeap []word

func (wh WordHeap) Len() int { return len(wh) }

func (wh WordHeap) Less(i, j int) bool {
	if wh[i].count == wh[j].count {
		return wh[i].word < wh[j].word
	}

	return wh[i].count > wh[j].count
}

func (wh WordHeap) Swap(i, j int) {
	wh[i], wh[j] = wh[j], wh[i]
}

func (wh *WordHeap) Push(x any) {
	item := x.(word)
	*wh = append(*wh, item)
}

func (wh *WordHeap) Pop() any {
	old := *wh
	n := len(old)
	item := old[n-1]
	*wh = old[0 : n-1]
	return item
}

func topKFrequent(words []string, k int) []string {
	wordCountMap := make(map[string]int, len(words))
	for _, word := range words {
		wordCountMap[word]++
	}

	wordHeap := WordHeap{}
	for w, count := range wordCountMap {
		wordItem := word{
			word:  w,
			count: count,
		}

		wordHeap = append(wordHeap, wordItem)
	}

	heap.Init(&wordHeap)

	var res []string
	for k > 0 {
		res = append(res, heap.Pop(&wordHeap).(word).word)
		k--
	}

	return res
}

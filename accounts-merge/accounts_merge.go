package accounts_merge

import "sort"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{
		parent: parent,
		rank:   rank,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = uf.parent[rootY]
		} else if uf.rank[rootX] > uf.parent[rootY] {
			uf.parent[rootY] = uf.parent[rootX]
		} else {
			uf.parent[rootY] = uf.parent[rootX]
			uf.rank[rootX]++
		}
	}
}

func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind(len(accounts))
	emails := make(map[string]int)

	for i, acc := range accounts {
		for _, email := range acc[1:] {
			if idx, ok := emails[email]; ok {
				uf.Union(i, idx)
			}

			emails[email] = i
		}
	}

	merged := make(map[int][]string)
	for email, idx := range emails {
		root := uf.Find(idx)
		merged[root] = append(merged[root], email)
	}

	var res [][]string
	for root, emails := range merged {
		sort.Strings(emails)
		name := accounts[root][0]
		account := append([]string{name}, emails...)
		res = append(res, account)
	}

	return res
}

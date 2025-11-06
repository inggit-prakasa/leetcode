package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(processQueries(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}})) // Example usage
}

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p, r := make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	return &DSU{p, r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) {
	pa, pb := d.Find(a), d.Find(b)
	if pa == pb {
		return
	}
	if d.rank[pa] < d.rank[pb] {
		pa, pb = pb, pa
	}
	d.parent[pb] = pa
	if d.rank[pa] == d.rank[pb] {
		d.rank[pa]++
	}
}

// --- min heap ---
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c)
	for _, conn := range connections {
		dsu.Union(conn[0], conn[1])
	}

	grids := make(map[int]*IntHeap)
	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		fmt.Println(i)
		fmt.Println(root)
		if grids[root] == nil {
			grids[root] = &IntHeap{}
		}
		fmt.Println(grids[root])
		heap.Push(grids[root], i)
	}

	offline := make([]bool, c+1)
	res := []int{}

	for _, q := range queries {
		typ, x := q[0], q[1]
		if typ == 2 {
			offline[x] = true
		} else {
			if !offline[x] {
				res = append(res, x)
				continue
			}
			root := dsu.Find(x)
			h := grids[root]
			for h.Len() > 0 && offline[(*h)[0]] {
				heap.Pop(h)
			}
			if h.Len() == 0 {
				res = append(res, -1)
			} else {
				res = append(res, (*h)[0])
			}
		}
	}
	return res
}

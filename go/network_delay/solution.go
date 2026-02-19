package network_delay

import "container/heap"

type edge struct {
	to, w int
}

type item struct {
	node, dist int
}

type minHeap []item

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)         { *h = append(*h, x.(item)) }
func (h *minHeap) Pop() any           { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make([][]edge, n+1)
	for _, t := range times {
		adj[t[0]] = append(adj[t[0]], edge{t[1], t[2]})
	}

	const inf = 1<<31 - 1
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0

	h := &minHeap{{k, 0}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range adj[cur.node] {
			if nd := cur.dist + e.w; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(h, item{e.to, nd})
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

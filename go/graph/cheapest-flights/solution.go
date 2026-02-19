package cheapest_flights

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1<<31 - 1
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0

	for i := 0; i <= k; i++ {
		prev := make([]int, n)
		copy(prev, dist)
		for _, f := range flights {
			from, to, price := f[0], f[1], f[2]
			if prev[from] != inf && prev[from]+price < dist[to] {
				dist[to] = prev[from] + price
			}
		}
	}

	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}

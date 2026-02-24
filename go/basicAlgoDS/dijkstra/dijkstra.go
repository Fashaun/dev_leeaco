package dijkstra

import (
	"fmt"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Adj map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{Adj: make(map[int][]Edge)}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.Adj[from] = append(g.Adj[from], Edge{To: to, Weight: weight})
	if _, ok := g.Adj[to]; !ok {
		g.Adj[to] = []Edge{}
	}
}

func Dijkstra(g *Graph, source int) map[int]int {
	dist := make(map[int]int)
	visited := make(map[int]bool)

	for v := range g.Adj {
		dist[v] = math.MaxInt
	}
	dist[source] = 0

	for i := 0; i < len(g.Adj); i++ {
		// Find unvisited vertex with minimum distance
		u := -1
		minDist := math.MaxInt
		for v := range g.Adj {
			if !visited[v] && dist[v] < minDist {
				u = v
				minDist = dist[v]
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true

		// Relax edges
		for _, edge := range g.Adj[u] {
			newDist := dist[u] + edge.Weight
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
			}
		}
	}

	return dist
}

func Run() {
	g := NewGraph()
	// Build undirected weighted graph (add edges in both directions)
	edges := [][3]int{
		{0, 1, 4},
		{0, 2, 1},
		{2, 1, 2},
		{1, 3, 1},
		{2, 3, 5},
		{3, 4, 3},
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1], e[2])
		g.AddEdge(e[1], e[0], e[2])
	}

	fmt.Println("=== Weighted Graph Edges ===")
	for from, neighbors := range g.Adj {
		for _, e := range neighbors {
			fmt.Printf("  %d -> %d (weight: %d)\n", from, e.To, e.Weight)
		}
	}

	source := 0
	dist := Dijkstra(g, source)

	fmt.Printf("\n=== Shortest Distances from Vertex %d ===\n", source)
	for v := 0; v < len(dist); v++ {
		if d, ok := dist[v]; ok {
			if d == math.MaxInt {
				fmt.Printf("  Vertex %d: unreachable\n", v)
			} else {
				fmt.Printf("  Vertex %d: %d\n", v, d)
			}
		}
	}
}

package topological_sort

import (
	"errors"
	"fmt"
)

type DAG struct {
	Adj map[int][]int
}

func NewDAG() *DAG {
	return &DAG{Adj: make(map[int][]int)}
}

func (g *DAG) AddEdge(from, to int) {
	g.Adj[from] = append(g.Adj[from], to)
	if _, ok := g.Adj[to]; !ok {
		g.Adj[to] = []int{}
	}
}

func TopologicalSort(g *DAG) ([]int, error) {
	// Compute in-degrees
	inDegree := make(map[int]int)
	for v := range g.Adj {
		if _, ok := inDegree[v]; !ok {
			inDegree[v] = 0
		}
		for _, neighbor := range g.Adj[v] {
			inDegree[neighbor]++
		}
	}

	// Enqueue vertices with in-degree 0
	queue := []int{}
	for v, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, v)
		}
	}

	result := []int{}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)

		for _, neighbor := range g.Adj[v] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != len(g.Adj) {
		return result, errors.New("graph contains a cycle")
	}

	return result, nil
}

func Run() {
	g := NewDAG()
	// Build a sample DAG
	edges := [][2]int{
		{5, 2},
		{5, 0},
		{4, 0},
		{4, 1},
		{2, 3},
		{3, 1},
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	fmt.Println("=== DAG Edges ===")
	for from, neighbors := range g.Adj {
		for _, to := range neighbors {
			fmt.Printf("  %d -> %d\n", from, to)
		}
	}

	order, err := TopologicalSort(g)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n=== Topological Order ===")
	fmt.Println(" ", order)
}

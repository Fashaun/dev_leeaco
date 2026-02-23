package graph_adjacency_list

import "fmt"

type Vertex struct {
	Val int
}

func NewVertex(val int) *Vertex {
	return &Vertex{Val: val}
}

type GraphAdjList struct {
	adjList map[*Vertex][]*Vertex
}

func NewGraphAdjList(edges [][]*Vertex) *GraphAdjList {
	g := &GraphAdjList{
		adjList: make(map[*Vertex][]*Vertex),
	}
	for _, edge := range edges {
		g.AddVertex(edge[0])
		g.AddVertex(edge[1])
		g.AddEdge(edge[0], edge[1])
	}
	return g
}

func (g *GraphAdjList) AddEdge(v1, v2 *Vertex) {
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}

func (g *GraphAdjList) RemoveEdge(v1, v2 *Vertex) {
	g.adjList[v1] = removeVertex(g.adjList[v1], v2)
	g.adjList[v2] = removeVertex(g.adjList[v2], v1)
}

func removeVertex(list []*Vertex, v *Vertex) []*Vertex {
	for i, vertex := range list {
		if vertex == v {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func (g *GraphAdjList) AddVertex(v *Vertex) {
	if _, ok := g.adjList[v]; ok {
		return
	}
	g.adjList[v] = make([]*Vertex, 0)
}

func (g *GraphAdjList) RemoveVertex(v *Vertex) {
	delete(g.adjList, v)
	for key := range g.adjList {
		g.adjList[key] = removeVertex(g.adjList[key], v)
	}
}

func (g *GraphAdjList) Print() {
	fmt.Println("Adjacency List:")
	for v, neighbors := range g.adjList {
		vals := make([]int, len(neighbors))
		for i, n := range neighbors {
			vals[i] = n.Val
		}
		fmt.Printf("  %d -> %v\n", v.Val, vals)
	}
}

func Run() {
	v0 := NewVertex(1)
	v1 := NewVertex(3)
	v2 := NewVertex(2)
	v3 := NewVertex(5)
	v4 := NewVertex(4)
	edges := [][]*Vertex{{v0, v1}, {v0, v3}, {v1, v2}, {v2, v3}, {v2, v4}, {v3, v4}}
	g := NewGraphAdjList(edges)

	fmt.Println("=== Initial Graph ===")
	g.Print()

	fmt.Println("\n=== Add Edge (v0, v2) ===")
	g.AddEdge(v0, v2)
	g.Print()

	fmt.Println("\n=== Remove Edge (v0, v1) ===")
	g.RemoveEdge(v0, v1)
	g.Print()

	v5 := NewVertex(6)
	fmt.Println("\n=== Add Vertex 6 ===")
	g.AddVertex(v5)
	g.Print()

	fmt.Println("\n=== Remove Vertex 3 (v1) ===")
	g.RemoveVertex(v1)
	g.Print()
}

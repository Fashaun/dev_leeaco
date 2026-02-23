package graph_adjacency_matrix

import "fmt"

type GraphAdjMat struct {
	vertices []int
	adjMat   [][]int
}

func NewGraphAdjMat(vertices []int, edges [][]int) *GraphAdjMat {
	n := len(vertices)
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	g := &GraphAdjMat{
		vertices: append([]int{}, vertices...),
		adjMat:   adjMat,
	}
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}
	return g
}

func (g *GraphAdjMat) AddEdge(i, j int) {
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

func (g *GraphAdjMat) RemoveEdge(i, j int) {
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

func (g *GraphAdjMat) AddVertex(val int) {
	n := len(g.vertices)
	g.vertices = append(g.vertices, val)
	newRow := make([]int, n+1)
	g.adjMat = append(g.adjMat, newRow)
	for i := 0; i < n; i++ {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

func (g *GraphAdjMat) RemoveVertex(index int) {
	n := len(g.vertices)
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	for i := 0; i < n-1; i++ {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

func (g *GraphAdjMat) Print() {
	fmt.Println("Vertices:", g.vertices)
	fmt.Println("Adjacency Matrix:")
	for _, row := range g.adjMat {
		fmt.Println(row)
	}
}

func Run() {
	vertices := []int{1, 3, 2, 5, 4}
	edges := [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 3}, {2, 4}, {3, 4}}
	g := NewGraphAdjMat(vertices, edges)

	fmt.Println("=== Initial Graph ===")
	g.Print()

	fmt.Println("\n=== Add Edge (0, 2) ===")
	g.AddEdge(0, 2)
	g.Print()

	fmt.Println("\n=== Remove Edge (0, 1) ===")
	g.RemoveEdge(0, 1)
	g.Print()

	fmt.Println("\n=== Add Vertex 6 ===")
	g.AddVertex(6)
	g.Print()

	fmt.Println("\n=== Remove Vertex at index 1 ===")
	g.RemoveVertex(1)
	g.Print()
}

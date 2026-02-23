## 1. Graph Package Setup

- [x] 1.1 Create `graph_adjacency_list/` directory
- [x] 1.2 Create `graph_adjacency_list/graph_adjacency_list.go` with `package graph_adjacency_list`

## 2. Core Data Structure

- [x] 2.1 Define `Vertex` struct with `Val int` field and `NewVertex(val int) *Vertex` constructor
- [x] 2.2 Define `GraphAdjList` struct with `adjList map[*Vertex][]*Vertex`
- [x] 2.3 Implement `NewGraphAdjList(edges [][]*Vertex) *GraphAdjList` constructor

## 3. Basic Operations

- [x] 3.1 Implement `AddEdge(v1, v2 *Vertex)` method
- [x] 3.2 Implement `RemoveEdge(v1, v2 *Vertex)` method
- [x] 3.3 Implement `AddVertex(v *Vertex)` method
- [x] 3.4 Implement `RemoveVertex(v *Vertex)` method
- [x] 3.5 Implement `Print()` method to display each vertex and its neighbor list

## 4. Demo Entry Point

- [x] 4.1 Implement `Run()` function that demonstrates all operations with printed output

## 5. CLI Registration

- [x] 5.1 Add `"basicAlgoDS/graph_adjacency_list"` import to `main.go`
- [x] 5.2 Add `case "graph_adjacency_list"` to switch in `main.go` calling `graph_adjacency_list.Run()`

## 6. Verification

- [x] 6.1 Run `go run main.go graph_adjacency_list` and verify graph operations demo output
- [x] 6.2 Run `go run main.go unknown` and verify error message still works

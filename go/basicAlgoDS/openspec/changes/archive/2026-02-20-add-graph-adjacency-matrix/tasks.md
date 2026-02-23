## 1. Graph Package Setup

- [x] 1.1 Create `graph_adjacency_matrix/` directory
- [x] 1.2 Create `graph_adjacency_matrix/graph_adjacency_matrix.go` with `package graph_adjacency_matrix`

## 2. Core Data Structure

- [x] 2.1 Define `GraphAdjMat` struct with `vertices []int` and `adjMat [][]int`
- [x] 2.2 Implement `NewGraphAdjMat(vertices []int, edges [][]int) *GraphAdjMat` constructor

## 3. Basic Operations

- [x] 3.1 Implement `AddEdge(i, j int)` method
- [x] 3.2 Implement `RemoveEdge(i, j int)` method
- [x] 3.3 Implement `AddVertex(val int)` method
- [x] 3.4 Implement `RemoveVertex(index int)` method
- [x] 3.5 Implement `Print()` method to display vertex list and adjacency matrix

## 4. Demo Entry Point

- [x] 4.1 Implement `Run()` function that demonstrates all operations with printed output

## 5. CLI Registration

- [x] 5.1 Add `"basicAlgoDS/graph_adjacency_matrix"` import to `main.go`
- [x] 5.2 Add `case "graph_adjacency_matrix"` to switch in `main.go` calling `graph_adjacency_matrix.Run()`

## 6. Verification

- [x] 6.1 Run `go run main.go graph_adjacency_matrix` and verify graph operations demo output
- [x] 6.2 Run `go run main.go unknown` and verify error message still works

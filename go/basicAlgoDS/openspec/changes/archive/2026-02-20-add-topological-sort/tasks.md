## 1. Create topological_sort package and data types

- [x] 1.1 Create `topological_sort/` directory and `topological_sort/topological_sort.go` with package declaration
- [x] 1.2 Define `DAG` struct with `Adj map[int][]int`, `NewDAG()` constructor, and `AddEdge(from, to int)` method

## 2. Implement topological sort algorithm

- [x] 2.1 Implement `TopologicalSort(g *DAG) ([]int, error)` using Kahn's algorithm with in-degree tracking and cycle detection

## 3. Run demo function

- [x] 3.1 Implement `Run()` function that builds a sample DAG, runs topological sort, and prints edges and ordering

## 4. Register in CLI dispatcher

- [x] 4.1 Add `topological_sort` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go topological_sort` and verify output

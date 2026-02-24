## 1. Create dijkstra package and data types

- [x] 1.1 Create `dijkstra/` directory and `dijkstra/dijkstra.go` with package declaration
- [x] 1.2 Define `Edge` struct with `To` and `Weight` fields
- [x] 1.3 Define `Graph` struct with `Adj map[int][]Edge`, `NewGraph()` constructor, and `AddEdge(from, to, weight int)` method

## 2. Implement Dijkstra algorithm

- [x] 2.1 Implement `Dijkstra(g *Graph, source int) map[int]int` using linear scan for min-distance selection and `math.MaxInt` for infinity

## 3. Run demo function

- [x] 3.1 Implement `Run()` function that builds a sample weighted graph, runs Dijkstra, and prints edges and shortest distances

## 4. Register in CLI dispatcher

- [x] 4.1 Add `dijkstra` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go dijkstra` and verify output

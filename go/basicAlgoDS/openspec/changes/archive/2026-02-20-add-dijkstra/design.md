## Context

The project has unweighted graph representations (adjacency list and matrix) but no graph algorithms. Dijkstra's algorithm requires a weighted graph, so this package needs its own lightweight weighted graph representation. The project uses `int` types throughout and keeps each example self-contained in a single package.

## Goals / Non-Goals

**Goals:**

- Implement Dijkstra's shortest path algorithm for weighted graphs with non-negative edge weights
- Include a simple weighted graph representation within the package
- Return shortest distances from a source vertex to all other vertices
- Match existing project conventions (single file, `Run()` demo, CLI registration)

**Non-Goals:**

- Reusing the existing `graph_adjacency_list` or `graph_adjacency_matrix` packages (they are unweighted)
- Returning the actual shortest paths (just distances)
- Priority queue via `container/heap` — keep it simple with a linear scan for minimum distance
- Negative edge weights or Bellman-Ford

## Decisions

### 1. Weighted graph: adjacency list via `map[int][]Edge`

Represent the graph as `map[int][]Edge` where `Edge` is a struct with `To` and `Weight` fields. Vertices are identified by `int` (not pointers like in `graph_adjacency_list`). This is simpler and sufficient for Dijkstra.

**Alternative considered:** Reuse `graph_adjacency_list.Vertex` pointers with added weight. Rejected — adds coupling and the existing types don't support weights.

### 2. Min-distance selection: linear scan over unvisited vertices

Each iteration, scan all vertices to find the unvisited one with minimum distance. This gives O(V^2) overall, which is the textbook approach and clearly demonstrates the algorithm.

**Alternative considered:** Binary heap / `container/heap` for O((V+E) log V). Rejected — adds significant complexity for an educational example.

### 3. Distance representation: `math.MaxInt` as infinity

Use `math.MaxInt` to represent unreachable vertices. The `Dijkstra` function returns `map[int]int` mapping each vertex to its shortest distance from the source.

### 4. Graph construction helper: `AddEdge(from, to, weight int)`

Provide a simple `AddEdge` method for building the graph. The graph is directed — for undirected graphs, the caller adds edges in both directions.

## Risks / Trade-offs

- [O(V^2) instead of O((V+E) log V)] → Acceptable for educational purposes. The linear scan approach is easier to understand.
- [Directed graph only] → The demo will add edges in both directions to simulate undirected. This is the standard approach and avoids hiding complexity.
- [No path reconstruction] → Returning only distances keeps the implementation focused. Path reconstruction can be added in a future change.

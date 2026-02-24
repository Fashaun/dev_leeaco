## Context

The project has graph data structures (adjacency list and matrix) and Dijkstra's algorithm, but no algorithm for ordering vertices in a DAG. Topological sort is a classic graph algorithm that produces a linear ordering of vertices such that for every directed edge (u, v), vertex u comes before v. This package needs its own directed graph representation since the existing graph packages don't fit (adjacency list uses pointer-based vertices, Dijkstra uses weighted edges).

## Goals / Non-Goals

**Goals:**

- Implement topological sort using Kahn's algorithm (BFS with in-degree tracking)
- Include a simple directed graph representation within the package
- Detect cycles and return an error if the graph is not a DAG
- Match existing project conventions (single file, `Run()` demo, CLI registration)

**Non-Goals:**

- DFS-based topological sort — Kahn's algorithm is more intuitive for educational purposes
- Reusing graph types from other packages
- Weighted edges or shortest paths
- All possible topological orderings (just one valid ordering)

## Decisions

### 1. Algorithm: Kahn's algorithm (BFS-based)

Use Kahn's algorithm: compute in-degrees, enqueue all vertices with in-degree 0, repeatedly dequeue a vertex, add to result, and decrement in-degrees of neighbors. This naturally produces a topological ordering and detects cycles (if result length < vertex count).

**Alternative considered:** DFS-based topological sort using post-order reversal. Rejected — Kahn's algorithm is more intuitive, directly shows the "peel off zero in-degree" logic, and naturally detects cycles.

### 2. Graph representation: `map[int][]int` adjacency list

Represent the directed graph as `map[int][]int` where keys are vertices and values are lists of neighbors. Vertices are `int`. Simpler than the Dijkstra package since no weights are needed.

### 3. Cycle detection: compare result length to vertex count

If Kahn's algorithm terminates with fewer vertices in the result than in the graph, the graph contains a cycle. Return the partial result and an error.

### 4. Return type: `([]int, error)`

`TopologicalSort` returns `([]int, error)` — the ordered vertices on success, or an error if a cycle is detected. This teaches Go's idiomatic error handling pattern.

## Risks / Trade-offs

- [Non-deterministic ordering for vertices with same in-degree] → Acceptable. The result is a valid topological order, though not necessarily unique. The demo will use a graph with a clear expected order.
- [Cycle detection returns error, not cycle location] → Keeps the implementation simple. Identifying the exact cycle vertices would add complexity without educational benefit.

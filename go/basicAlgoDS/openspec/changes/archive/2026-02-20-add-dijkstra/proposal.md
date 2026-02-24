## Why

The project has graph data structures (adjacency list and matrix) but no graph algorithms. Dijkstra's algorithm is the foundational shortest-path algorithm for weighted graphs with non-negative edges, and a key topic in algorithm education.

## What Changes

- Add new `dijkstra` package with a weighted graph representation and `Dijkstra(graph, source)` function that computes shortest distances from a source vertex to all other vertices
- Add `Run()` demo function that builds a sample weighted graph, runs Dijkstra, and prints shortest distances
- Register `dijkstra` in the CLI dispatcher in `main.go`

## Capabilities

### New Capabilities

- `dijkstra`: Dijkstra's shortest path algorithm with weighted graph representation, distance computation, and demo output

### Modified Capabilities

- `cli-entry`: Add `dijkstra` case to the CLI switch dispatcher

## Impact

- New file: `dijkstra/dijkstra.go`
- Modified file: `main.go` (new import and switch case)

## Why

The project has graph data structures and Dijkstra's shortest path but is missing topological sorting — a fundamental algorithm for directed acyclic graphs (DAGs). Topological sort is essential for dependency resolution, task scheduling, and build systems, making it a key topic in algorithm education.

## What Changes

- Add new `topological_sort` package with a DAG representation and `TopologicalSort(graph)` function that returns vertices in topological order using Kahn's algorithm (BFS-based)
- Add `Run()` demo function that builds a sample DAG, runs topological sort, and prints the ordering
- Register `topological_sort` in the CLI dispatcher in `main.go`

## Capabilities

### New Capabilities

- `topological-sort`: Topological sorting algorithm for DAGs with directed graph representation, cycle detection, and demo output

### Modified Capabilities

- `cli-entry`: Add `topological_sort` case to the CLI switch dispatcher

## Impact

- New file: `topological_sort/topological_sort.go`
- Modified file: `main.go` (new import and switch case)

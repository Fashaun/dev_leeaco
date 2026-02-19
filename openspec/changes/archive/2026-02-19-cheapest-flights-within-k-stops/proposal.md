## Why

The LeetCode solutions repo currently covers hash maps (Two Sum), sorting (3Sum), and graph shortest-path (Network Delay Time / Dijkstra). Adding problem #787 "Cheapest Flights Within K Stops" extends the graph category with a constrained shortest-path problem that requires a different algorithm (Bellman-Ford or BFS with pruning) due to the hop limit.

## What Changes

- Add a new solution directory `go/graph/cheapest-flights/` following the existing project structure
- Implement the `findCheapestPrice` function using a Bellman-Ford approach bounded by k+1 iterations
- Add table-driven tests covering all provided examples plus edge cases
- Add problem description and CLAUDE.md for the new problem

## Capabilities

### New Capabilities
- `cheapest-flights-solution`: Go implementation of LeetCode #787 with Bellman-Ford algorithm, bounded by at most k stops

### Modified Capabilities
<!-- None — this is a new standalone problem -->

## Impact

- `go/graph/cheapest-flights/solution.go`: New solution file
- `go/graph/cheapest-flights/solution_test.go`: New test file
- `go/graph/cheapest-flights/go.mod`: New module
- `go/graph/cheapest-flights/CheapestFlights.md`: Problem description
- `go/graph/cheapest-flights/CLAUDE.md`: Build/test instructions

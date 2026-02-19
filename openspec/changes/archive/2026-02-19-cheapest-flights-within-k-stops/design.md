## Context

The repo already has a graph shortest-path solution (Network Delay Time) using Dijkstra's algorithm with a min-heap. Problem #787 adds a constraint—at most k stops—which means standard Dijkstra doesn't directly apply because it optimizes for shortest distance without hop limits. We need an algorithm that naturally respects the stop constraint.

## Goals / Non-Goals

**Goals:**
- Solve LeetCode #787 correctly for all constraints (n <= 100, prices <= 10^4)
- Use an algorithm that clearly enforces the k-stop limit
- Follow existing repo conventions (package per problem, table-driven tests, CLAUDE.md)

**Non-Goals:**
- Optimizing for very large graphs (constraints are small: n <= 100)
- Supporting generic k-constrained shortest path as a reusable library
- Modifying any existing solutions

## Decisions

### Decision 1: Bellman-Ford over modified Dijkstra or BFS

**Choice:** Bellman-Ford with k+1 relaxation rounds.

**Why:** Bellman-Ford naturally limits the number of edges used. Running exactly k+1 iterations guarantees paths use at most k+1 edges (k intermediate stops). The implementation is straightforward—no priority queue or visited-state tracking needed.

**Alternatives considered:**
- **Modified Dijkstra with (node, stops) state:** Works but more complex—requires a priority queue with composite state and careful pruning. Overkill for n <= 100.
- **BFS level-by-level:** Viable but less clean for weighted graphs. Would need to track minimum cost per node per level.

**Trade-off:** Bellman-Ford is O(k * E) which is fine for the given constraints. It's simpler to implement and reason about correctness.

### Decision 2: Copy dist array each iteration

Each Bellman-Ford iteration must read from the *previous* iteration's distances (not the current one being updated). This prevents using paths with more edges than intended in a single iteration. A temporary copy of the dist array handles this.

### Decision 3: Directory structure under `go/graph/`

Place the solution at `go/graph/cheapest-flights/` to group it with other graph problems. The existing `network_delay` is at `go/network_delay/` (top-level), but going forward `go/graph/` is a better organizational pattern.

## Risks / Trade-offs

- **Bellman-Ford is O(k * E) vs Dijkstra's O(E log V):** Not a concern given n <= 100, but worth noting for documentation purposes.
- **New `go/graph/` directory diverges from `go/network_delay/` placement:** Minor inconsistency, but establishes better organization going forward. Not worth moving existing code.

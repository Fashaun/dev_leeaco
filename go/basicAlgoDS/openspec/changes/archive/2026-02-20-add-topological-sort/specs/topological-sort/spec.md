## ADDED Requirements

### Requirement: DAG struct with AddEdge

The `topological_sort` package SHALL provide a `DAG` struct backed by `map[int][]int` and a `NewDAG()` constructor. The `AddEdge(from, to int)` method SHALL add a directed edge from vertex `from` to vertex `to`.

#### Scenario: Add a directed edge

- **WHEN** `AddEdge(0, 1)` is called
- **THEN** the graph SHALL contain a directed edge from vertex 0 to vertex 1

#### Scenario: Add vertex with no outgoing edges

- **WHEN** `AddEdge(0, 1)` is called and vertex 1 has no outgoing edges
- **THEN** vertex 1 SHALL exist in the graph with an empty neighbor list

### Requirement: TopologicalSort function returns vertices in topological order

The `topological_sort` package SHALL provide a `TopologicalSort(g *DAG) ([]int, error)` function that returns a valid topological ordering of the DAG's vertices using Kahn's algorithm. If the graph contains a cycle, the function SHALL return an error.

#### Scenario: Sort a valid DAG

- **WHEN** `TopologicalSort` is called on a DAG with edges 5→2, 5→0, 4→0, 4→1, 2→3, 3→1
- **THEN** the returned slice SHALL contain all 6 vertices in an order where every edge (u, v) has u before v

#### Scenario: Source vertex has distance zero in-degree

- **WHEN** `TopologicalSort` is called on a valid DAG
- **THEN** vertices with in-degree 0 SHALL appear before their dependents

#### Scenario: Single vertex graph

- **WHEN** `TopologicalSort` is called on a DAG with only vertex 0
- **THEN** the result SHALL be `[0]` with no error

#### Scenario: Graph with a cycle

- **WHEN** `TopologicalSort` is called on a graph containing a cycle (e.g., 0→1, 1→2, 2→0)
- **THEN** the function SHALL return an error indicating a cycle was detected

#### Scenario: Empty graph

- **WHEN** `TopologicalSort` is called on a DAG with no vertices
- **THEN** the result SHALL be an empty slice with no error

### Requirement: Run function displays topological sort demo

The `topological_sort` package SHALL export a `Run()` function that builds a sample DAG, runs topological sort, and prints the edges and the resulting topological order.

#### Scenario: Run is called

- **WHEN** `topological_sort.Run()` is invoked
- **THEN** the function SHALL print the DAG edges and the topological ordering of vertices

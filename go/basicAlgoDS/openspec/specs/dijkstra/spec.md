### Requirement: Edge struct

The `dijkstra` package SHALL provide an `Edge` struct with `To int` and `Weight int` fields representing a directed weighted edge.

#### Scenario: Edge represents a weighted connection

- **WHEN** an `Edge{To: 2, Weight: 5}` is created
- **THEN** it SHALL represent a directed edge to vertex 2 with weight 5

### Requirement: Graph struct with AddEdge

The `dijkstra` package SHALL provide a `Graph` struct backed by `map[int][]Edge` and a `NewGraph()` constructor. The `AddEdge(from, to, weight int)` method SHALL add a directed weighted edge.

#### Scenario: Add a directed edge

- **WHEN** `AddEdge(0, 1, 4)` is called
- **THEN** the graph SHALL contain a directed edge from vertex 0 to vertex 1 with weight 4

#### Scenario: Add edges in both directions for undirected behavior

- **WHEN** `AddEdge(0, 1, 4)` and `AddEdge(1, 0, 4)` are both called
- **THEN** the graph SHALL contain edges in both directions between vertex 0 and vertex 1

### Requirement: Dijkstra function computes shortest distances

The `dijkstra` package SHALL provide a `Dijkstra(g *Graph, source int) map[int]int` function that computes the shortest distance from the source vertex to all other vertices using Dijkstra's algorithm. Unreachable vertices SHALL have distance `math.MaxInt`.

#### Scenario: Shortest distances from a source vertex

- **WHEN** `Dijkstra` is called on a graph with vertices 0-4 and weighted edges, with source 0
- **THEN** the returned map SHALL contain the shortest distance from vertex 0 to every other vertex

#### Scenario: Source vertex has distance zero

- **WHEN** `Dijkstra` is called with source vertex 0
- **THEN** the distance to vertex 0 SHALL be 0

#### Scenario: Unreachable vertex

- **WHEN** a vertex has no path from the source
- **THEN** the distance SHALL be `math.MaxInt`

#### Scenario: Single vertex graph

- **WHEN** `Dijkstra` is called on a graph with only one vertex as the source
- **THEN** the distance to that vertex SHALL be 0

### Requirement: Run function displays shortest path demo

The `dijkstra` package SHALL export a `Run()` function that builds a sample weighted graph, runs Dijkstra from a source vertex, and prints the shortest distances.

#### Scenario: Run is called

- **WHEN** `dijkstra.Run()` is invoked
- **THEN** the function SHALL print the graph edges and the shortest distances from the source vertex

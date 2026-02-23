### Requirement: Vertex struct

The `graph_adjacency_list` package SHALL export a `Vertex` struct with a `Val int` field, and a `NewVertex(val int) *Vertex` constructor function.

#### Scenario: Create a vertex

- **WHEN** `NewVertex(1)` is called
- **THEN** a `*Vertex` SHALL be returned with `Val == 1`

### Requirement: Graph initialization with adjacency list

The `graph_adjacency_list` package SHALL export a `NewGraphAdjList(edges [][]*Vertex)` function that creates a `*GraphAdjList`. Each edge is represented as `[]*Vertex{v1, v2}`. The function SHALL register all vertices and their adjacency relationships.

#### Scenario: Create graph from edges

- **WHEN** `NewGraphAdjList([][]*Vertex{{v0, v1}, {v0, v3}, {v1, v2}, {v2, v3}})` is called
- **THEN** a `*GraphAdjList` SHALL be returned with all vertices registered and each vertex's neighbor list populated symmetrically (undirected)

### Requirement: Add edge

The `GraphAdjList` SHALL provide an `AddEdge(v1, v2 *Vertex)` method that adds each vertex to the other's neighbor list.

#### Scenario: Add an edge between two vertices

- **WHEN** `AddEdge(v0, v2)` is called where v0 and v2 are not neighbors
- **THEN** v2 SHALL appear in v0's neighbor list and v0 SHALL appear in v2's neighbor list

### Requirement: Remove edge

The `GraphAdjList` SHALL provide a `RemoveEdge(v1, v2 *Vertex)` method that removes each vertex from the other's neighbor list.

#### Scenario: Remove an edge between two vertices

- **WHEN** `RemoveEdge(v0, v1)` is called where v0 and v1 are neighbors
- **THEN** v1 SHALL be removed from v0's neighbor list and v0 SHALL be removed from v1's neighbor list

### Requirement: Add vertex

The `GraphAdjList` SHALL provide an `AddVertex(v *Vertex)` method that registers a new vertex with an empty neighbor list.

#### Scenario: Add a new vertex

- **WHEN** `AddVertex(v)` is called with a new vertex
- **THEN** the vertex SHALL be registered in the adjacency list with an empty neighbor list

### Requirement: Remove vertex

The `GraphAdjList` SHALL provide a `RemoveVertex(v *Vertex)` method that removes the vertex and all edges connected to it from the adjacency list.

#### Scenario: Remove a vertex

- **WHEN** `RemoveVertex(v1)` is called
- **THEN** v1 SHALL be removed from the adjacency list and SHALL be removed from all other vertices' neighbor lists

### Requirement: Print graph

The `GraphAdjList` SHALL provide a `Print()` method that prints each vertex and its neighbor list to stdout.

#### Scenario: Print the graph

- **WHEN** `Print()` is called on a graph
- **THEN** each vertex and its neighbors SHALL be printed to stdout

### Requirement: Run function displays graph operations demo

The `graph_adjacency_list` package SHALL export a `Run()` function that demonstrates graph initialization, adding/removing edges, adding/removing vertices, and printing the graph.

#### Scenario: Run is called

- **WHEN** `graph_adjacency_list.Run()` is invoked
- **THEN** the function SHALL demonstrate all basic graph operations with printed output

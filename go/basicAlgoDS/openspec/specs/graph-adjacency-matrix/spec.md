### Requirement: Graph initialization with adjacency matrix

The `graph_adjacency_matrix` package SHALL export a `NewGraphAdjMat(vertices []int, edges [][]int)` function that creates a `*GraphAdjMat` with the given vertices and edges. Each edge is represented as `[]int{i, j}` where `i` and `j` are indices into the vertices slice. The adjacency matrix SHALL be symmetric (undirected graph) using 0/1 values.

#### Scenario: Create graph with vertices and edges

- **WHEN** `NewGraphAdjMat([]int{1, 3, 2, 5, 4}, [][]int{{0, 1}, {0, 3}, {1, 2}, {2, 3}, {2, 4}, {3, 4}})` is called
- **THEN** a `*GraphAdjMat` SHALL be returned with a 5x5 adjacency matrix where `adjMat[i][j] == 1` and `adjMat[j][i] == 1` for each edge `{i, j}`

### Requirement: Add edge

The `GraphAdjMat` SHALL provide an `AddEdge(i, j int)` method that sets `adjMat[i][j]` and `adjMat[j][i]` to 1.

#### Scenario: Add an edge between two vertices

- **WHEN** `AddEdge(0, 2)` is called on a graph where vertices at index 0 and 2 are not connected
- **THEN** `adjMat[0][2]` and `adjMat[2][0]` SHALL both be set to 1

### Requirement: Remove edge

The `GraphAdjMat` SHALL provide a `RemoveEdge(i, j int)` method that sets `adjMat[i][j]` and `adjMat[j][i]` to 0.

#### Scenario: Remove an edge between two vertices

- **WHEN** `RemoveEdge(0, 1)` is called on a graph where vertices at index 0 and 1 are connected
- **THEN** `adjMat[0][1]` and `adjMat[1][0]` SHALL both be set to 0

### Requirement: Add vertex

The `GraphAdjMat` SHALL provide an `AddVertex(val int)` method that appends a new vertex to the vertices list and expands the adjacency matrix by one row and one column (initialized to 0).

#### Scenario: Add a new vertex

- **WHEN** `AddVertex(6)` is called on a graph with 5 vertices
- **THEN** the vertices list SHALL contain 6 elements and the adjacency matrix SHALL be 6x6

### Requirement: Remove vertex

The `GraphAdjMat` SHALL provide a `RemoveVertex(index int)` method that removes the vertex at the given index from the vertices list and removes the corresponding row and column from the adjacency matrix.

#### Scenario: Remove a vertex by index

- **WHEN** `RemoveVertex(1)` is called on a graph with 5 vertices
- **THEN** the vertices list SHALL contain 4 elements, the adjacency matrix SHALL be 4x4, and all edges connected to the removed vertex SHALL be gone

### Requirement: Print graph

The `GraphAdjMat` SHALL provide a `Print()` method that prints the vertex list and the adjacency matrix to stdout.

#### Scenario: Print the graph

- **WHEN** `Print()` is called on a graph
- **THEN** the vertex list and adjacency matrix SHALL be printed to stdout

### Requirement: Run function displays graph operations demo

The `graph_adjacency_matrix` package SHALL export a `Run()` function that demonstrates graph initialization, adding/removing edges, adding/removing vertices, and printing the graph.

#### Scenario: Run is called

- **WHEN** `graph_adjacency_matrix.Run()` is invoked
- **THEN** the function SHALL demonstrate all basic graph operations with printed output

## MODIFIED Requirements

### Requirement: CLI argument dispatches to example package

The system SHALL read `os.Args[1]` as the package/folder name and invoke the corresponding example package's `Run()` function.

#### Scenario: Valid package name provided

- **WHEN** user runs `go run main.go go_algods_example`
- **THEN** the system SHALL call `go_algods_example.Run()` and display the package's output

#### Scenario: Bubble sort package name provided

- **WHEN** user runs `go run main.go bubble_sort`
- **THEN** the system SHALL call `bubble_sort.Run()` and display the sorting demo output

#### Scenario: Graph adjacency matrix package name provided

- **WHEN** user runs `go run main.go graph_adjacency_matrix`
- **THEN** the system SHALL call `graph_adjacency_matrix.Run()` and display the graph operations demo output

#### Scenario: Graph adjacency list package name provided

- **WHEN** user runs `go run main.go graph_adjacency_list`
- **THEN** the system SHALL call `graph_adjacency_list.Run()` and display the graph operations demo output

#### Scenario: Unknown package name provided

- **WHEN** user provides a package name that does not match any registered example
- **THEN** the system SHALL print an error message indicating the package is not found

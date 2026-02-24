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

#### Scenario: Stack linked list package name provided

- **WHEN** user runs `go run main.go stack_linked_list`
- **THEN** the system SHALL call `stack_linked_list.Run()` and display the stack operations demo output

#### Scenario: Stack array package name provided

- **WHEN** user runs `go run main.go stack_array`
- **THEN** the system SHALL call `stack_array.Run()` and display the stack operations demo output

#### Scenario: Queue linked list package name provided

- **WHEN** user runs `go run main.go queue_linked_list`
- **THEN** the system SHALL call `queue_linked_list.Run()` and display the queue operations demo output

#### Scenario: Queue array package name provided

- **WHEN** user runs `go run main.go queue_array`
- **THEN** the system SHALL call `queue_array.Run()` and display the queue operations demo output

#### Scenario: Deque array package name provided

- **WHEN** user runs `go run main.go deque_array`
- **THEN** the system SHALL call `deque_array.Run()` and display the deque operations demo output

#### Scenario: Selection sort package name provided

- **WHEN** user runs `go run main.go selection_sort`
- **THEN** the system SHALL call `selection_sort.Run()` and display the sorting demo output

#### Scenario: Insertion sort package name provided

- **WHEN** user runs `go run main.go insertion_sort`
- **THEN** the system SHALL call `insertion_sort.Run()` and display the sorting demo output

#### Scenario: Dijkstra package name provided

- **WHEN** user runs `go run main.go dijkstra`
- **THEN** the system SHALL call `dijkstra.Run()` and display the shortest path demo output

#### Scenario: Topological sort package name provided

- **WHEN** user runs `go run main.go topological_sort`
- **THEN** the system SHALL call `topological_sort.Run()` and display the topological ordering demo output

#### Scenario: BST operations package name provided

- **WHEN** user runs `go run main.go bst_operations`
- **THEN** the system SHALL call `bst_operations.Run()` and display the BST operations demo output

#### Scenario: Merge sort package name provided

- **WHEN** user runs `go run main.go merge_sort`
- **THEN** the system SHALL call `merge_sort.Run()` and display the sorting demo output

#### Scenario: Unknown package name provided

- **WHEN** user provides a package name that does not match any registered example
- **THEN** the system SHALL print an error message indicating the package is not found

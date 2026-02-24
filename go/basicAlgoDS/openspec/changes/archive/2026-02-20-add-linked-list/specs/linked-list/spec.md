## ADDED Requirements

### Requirement: ListNode struct and LinkedList container

The `linked_list` package SHALL provide a `ListNode` struct with `Val int` and `Next *ListNode` fields, a `NewListNode(val int) *ListNode` constructor, and a `LinkedList` struct with `Head *ListNode` field and a `NewLinkedList() *LinkedList` constructor.

#### Scenario: Create a new list node

- **WHEN** `NewListNode(5)` is called
- **THEN** a node with `Val=5` and `Next=nil` SHALL be returned

#### Scenario: Create a new empty linked list

- **WHEN** `NewLinkedList()` is called
- **THEN** a linked list with `Head=nil` SHALL be returned

### Requirement: InsertAtIndex inserts a node at a given position

The `linked_list` package SHALL provide an `InsertAtIndex(index int, val int)` method on `*LinkedList` that inserts a new node with the given value before the element at `index`.

#### Scenario: Insert at the head (index 0)

- **WHEN** `InsertAtIndex(0, 10)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[10, 1, 2, 3]`

#### Scenario: Insert in the middle

- **WHEN** `InsertAtIndex(2, 10)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[1, 2, 10, 3]`

#### Scenario: Insert at the tail

- **WHEN** `InsertAtIndex(3, 10)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[1, 2, 3, 10]`

#### Scenario: Insert into an empty list

- **WHEN** `InsertAtIndex(0, 1)` is called on an empty list
- **THEN** the list SHALL become `[1]`

### Requirement: RemoveAtIndex removes a node at a given position

The `linked_list` package SHALL provide a `RemoveAtIndex(index int)` method on `*LinkedList` that removes the node at the given index.

#### Scenario: Remove the head (index 0)

- **WHEN** `RemoveAtIndex(0)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[2, 3]`

#### Scenario: Remove from the middle

- **WHEN** `RemoveAtIndex(1)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[1, 3]`

#### Scenario: Remove the tail

- **WHEN** `RemoveAtIndex(2)` is called on a list `[1, 2, 3]`
- **THEN** the list SHALL become `[1, 2]`

### Requirement: Access retrieves a node's value by index

The `linked_list` package SHALL provide an `Access(index int) int` method on `*LinkedList` that returns the value of the node at the given index.

#### Scenario: Access the head

- **WHEN** `Access(0)` is called on a list `[1, 2, 3]`
- **THEN** the method SHALL return `1`

#### Scenario: Access a middle element

- **WHEN** `Access(1)` is called on a list `[1, 2, 3]`
- **THEN** the method SHALL return `2`

### Requirement: Find searches for a value and returns its index

The `linked_list` package SHALL provide a `Find(val int) int` method on `*LinkedList` that returns the index of the first node with the given value, or `-1` if not found.

#### Scenario: Find an existing value

- **WHEN** `Find(3)` is called on a list `[1, 2, 3, 4, 5]`
- **THEN** the method SHALL return `2`

#### Scenario: Find a non-existing value

- **WHEN** `Find(9)` is called on a list `[1, 2, 3, 4, 5]`
- **THEN** the method SHALL return `-1`

### Requirement: Print displays the linked list

The `linked_list` package SHALL provide a `Print()` method on `*LinkedList` that prints all values in the format `val1 -> val2 -> ... -> nil`.

#### Scenario: Print a non-empty list

- **WHEN** `Print()` is called on a list `[1, 2, 3]`
- **THEN** the output SHALL be `1 -> 2 -> 3 -> nil`

#### Scenario: Print an empty list

- **WHEN** `Print()` is called on an empty list
- **THEN** the output SHALL be `nil`

### Requirement: Run function displays linked list operations demo

The `linked_list` package SHALL export a `Run()` function that demonstrates initializing, inserting, removing, accessing, finding, and printing a singly linked list.

#### Scenario: Run is called

- **WHEN** `linked_list.Run()` is invoked
- **THEN** the function SHALL print the list state after each operation demonstrating all 5 core operations

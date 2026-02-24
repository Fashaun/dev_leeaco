### Requirement: Node struct

The `queue_linked_list` package SHALL define a `Node` struct with `Val int` and `Next *Node` fields.

#### Scenario: Node holds value and pointer

- **WHEN** a `Node` is created with `Val: 5`
- **THEN** the node SHALL have `Val == 5` and `Next == nil`

### Requirement: LinkedListQueue struct

The `queue_linked_list` package SHALL define a `LinkedListQueue` struct with `front *Node`, `rear *Node`, and `size int` fields.

#### Scenario: LinkedListQueue initialization

- **WHEN** a `LinkedListQueue` is created
- **THEN** it SHALL have `front == nil`, `rear == nil`, and `Size()` SHALL return 0

### Requirement: Enqueue operation

The `LinkedListQueue` SHALL provide an `Enqueue(val int)` method that inserts a new node at the rear of the queue.

#### Scenario: Enqueue a value onto an empty queue

- **WHEN** `Enqueue(5)` is called on an empty queue
- **THEN** `Peek()` SHALL return `(5, true)` and `Size()` SHALL return 1

#### Scenario: Enqueue multiple values

- **WHEN** `Enqueue(1)`, `Enqueue(2)`, `Enqueue(3)` are called sequentially
- **THEN** `Peek()` SHALL return `(1, true)` and `Size()` SHALL return 3

### Requirement: Dequeue operation

The `LinkedListQueue` SHALL provide a `Dequeue() (int, bool)` method that removes and returns the front element.

#### Scenario: Dequeue from a non-empty queue

- **WHEN** `Dequeue()` is called on a queue with front value 1
- **THEN** it SHALL return `(1, true)` and the queue size SHALL decrease by 1

#### Scenario: Dequeue from an empty queue

- **WHEN** `Dequeue()` is called on an empty queue
- **THEN** it SHALL return `(0, false)`

### Requirement: Peek operation

The `LinkedListQueue` SHALL provide a `Peek() (int, bool)` method that returns the front element without removing it.

#### Scenario: Peek at a non-empty queue

- **WHEN** `Peek()` is called on a queue with front value 1
- **THEN** it SHALL return `(1, true)` and the queue size SHALL remain unchanged

#### Scenario: Peek at an empty queue

- **WHEN** `Peek()` is called on an empty queue
- **THEN** it SHALL return `(0, false)`

### Requirement: Size operation

The `LinkedListQueue` SHALL provide a `Size() int` method that returns the number of elements in the queue.

#### Scenario: Size of queue

- **WHEN** `Size()` is called on a queue with 3 elements
- **THEN** it SHALL return 3

### Requirement: IsEmpty operation

The `LinkedListQueue` SHALL provide an `IsEmpty() bool` method that returns true when the queue has no elements.

#### Scenario: Empty queue

- **WHEN** `IsEmpty()` is called on an empty queue
- **THEN** it SHALL return true

#### Scenario: Non-empty queue

- **WHEN** `IsEmpty()` is called on a queue with elements
- **THEN** it SHALL return false

### Requirement: Print operation

The `LinkedListQueue` SHALL provide a `Print()` method that prints all queue elements from front to rear to stdout.

#### Scenario: Print queue contents

- **WHEN** `Print()` is called on a queue containing [1, 2, 3] (front to rear)
- **THEN** the elements SHALL be printed to stdout in front-to-rear order

### Requirement: Run function displays queue operations demo

The `queue_linked_list` package SHALL export a `Run()` function that demonstrates Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations.

#### Scenario: Run is called

- **WHEN** `queue_linked_list.Run()` is invoked
- **THEN** the function SHALL demonstrate all queue operations with printed output

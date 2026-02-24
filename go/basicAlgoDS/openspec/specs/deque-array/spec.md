### Requirement: ArrayDeque struct

The `deque_array` package SHALL define an `ArrayDeque` struct with an internal `[]int` slice and a `size int` field to track the number of elements.

#### Scenario: ArrayDeque initialization

- **WHEN** an `ArrayDeque` is created
- **THEN** it SHALL have an empty slice and `Size()` SHALL return 0

### Requirement: PushFront operation

The `ArrayDeque` SHALL provide a `PushFront(val int)` method that inserts a new element at the front of the deque.

#### Scenario: PushFront on an empty deque

- **WHEN** `PushFront(5)` is called on an empty deque
- **THEN** `PeekFront()` SHALL return `(5, true)` and `Size()` SHALL return 1

#### Scenario: PushFront on a non-empty deque

- **WHEN** `PushFront(0)` is called on a deque containing [1, 2, 3]
- **THEN** `PeekFront()` SHALL return `(0, true)` and `Size()` SHALL return 4

### Requirement: PushRear operation

The `ArrayDeque` SHALL provide a `PushRear(val int)` method that inserts a new element at the rear of the deque.

#### Scenario: PushRear on an empty deque

- **WHEN** `PushRear(5)` is called on an empty deque
- **THEN** `PeekRear()` SHALL return `(5, true)` and `Size()` SHALL return 1

#### Scenario: PushRear on a non-empty deque

- **WHEN** `PushRear(4)` is called on a deque containing [1, 2, 3]
- **THEN** `PeekRear()` SHALL return `(4, true)` and `Size()` SHALL return 4

### Requirement: PopFront operation

The `ArrayDeque` SHALL provide a `PopFront() (int, bool)` method that removes and returns the front element.

#### Scenario: PopFront from a non-empty deque

- **WHEN** `PopFront()` is called on a deque with front value 1
- **THEN** it SHALL return `(1, true)` and the deque size SHALL decrease by 1

#### Scenario: PopFront from an empty deque

- **WHEN** `PopFront()` is called on an empty deque
- **THEN** it SHALL return `(0, false)`

### Requirement: PopRear operation

The `ArrayDeque` SHALL provide a `PopRear() (int, bool)` method that removes and returns the rear element.

#### Scenario: PopRear from a non-empty deque

- **WHEN** `PopRear()` is called on a deque with rear value 3
- **THEN** it SHALL return `(3, true)` and the deque size SHALL decrease by 1

#### Scenario: PopRear from an empty deque

- **WHEN** `PopRear()` is called on an empty deque
- **THEN** it SHALL return `(0, false)`

### Requirement: PeekFront operation

The `ArrayDeque` SHALL provide a `PeekFront() (int, bool)` method that returns the front element without removing it.

#### Scenario: PeekFront at a non-empty deque

- **WHEN** `PeekFront()` is called on a deque with front value 1
- **THEN** it SHALL return `(1, true)` and the deque size SHALL remain unchanged

#### Scenario: PeekFront at an empty deque

- **WHEN** `PeekFront()` is called on an empty deque
- **THEN** it SHALL return `(0, false)`

### Requirement: PeekRear operation

The `ArrayDeque` SHALL provide a `PeekRear() (int, bool)` method that returns the rear element without removing it.

#### Scenario: PeekRear at a non-empty deque

- **WHEN** `PeekRear()` is called on a deque with rear value 3
- **THEN** it SHALL return `(3, true)` and the deque size SHALL remain unchanged

#### Scenario: PeekRear at an empty deque

- **WHEN** `PeekRear()` is called on an empty deque
- **THEN** it SHALL return `(0, false)`

### Requirement: Size operation

The `ArrayDeque` SHALL provide a `Size() int` method that returns the number of elements in the deque.

#### Scenario: Size of deque

- **WHEN** `Size()` is called on a deque with 3 elements
- **THEN** it SHALL return 3

### Requirement: IsEmpty operation

The `ArrayDeque` SHALL provide an `IsEmpty() bool` method that returns true when the deque has no elements.

#### Scenario: Empty deque

- **WHEN** `IsEmpty()` is called on an empty deque
- **THEN** it SHALL return true

#### Scenario: Non-empty deque

- **WHEN** `IsEmpty()` is called on a deque with elements
- **THEN** it SHALL return false

### Requirement: Print operation

The `ArrayDeque` SHALL provide a `Print()` method that prints all deque elements from front to rear to stdout.

#### Scenario: Print deque contents

- **WHEN** `Print()` is called on a deque containing [1, 2, 3] (front to rear)
- **THEN** the elements SHALL be printed to stdout in front-to-rear order

### Requirement: Run function displays deque operations demo

The `deque_array` package SHALL export a `Run()` function that demonstrates PushFront, PushRear, PopFront, PopRear, PeekFront, PeekRear, Size, IsEmpty, and Print operations.

#### Scenario: Run is called

- **WHEN** `deque_array.Run()` is invoked
- **THEN** the function SHALL demonstrate all deque operations with printed output

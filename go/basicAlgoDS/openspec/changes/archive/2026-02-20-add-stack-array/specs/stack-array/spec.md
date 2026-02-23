## ADDED Requirements

### Requirement: ArrayStack struct

The `stack_array` package SHALL define an `ArrayStack` struct with an internal `[]int` slice and a `size int` field to track the number of elements.

#### Scenario: ArrayStack initialization

- **WHEN** an `ArrayStack` is created
- **THEN** it SHALL have an empty slice and `Size()` SHALL return 0

### Requirement: Push operation

The `ArrayStack` SHALL provide a `Push(val int)` method that appends a new element to the top of the stack.

#### Scenario: Push a value onto the stack

- **WHEN** `Push(5)` is called on an empty stack
- **THEN** `Peek()` SHALL return `(5, true)` and `Size()` SHALL return 1

#### Scenario: Push multiple values

- **WHEN** `Push(1)`, `Push(2)`, `Push(3)` are called sequentially
- **THEN** `Peek()` SHALL return `(3, true)` and `Size()` SHALL return 3

### Requirement: Pop operation

The `ArrayStack` SHALL provide a `Pop() (int, bool)` method that removes and returns the top element.

#### Scenario: Pop from a non-empty stack

- **WHEN** `Pop()` is called on a stack with top value 3
- **THEN** it SHALL return `(3, true)` and the stack size SHALL decrease by 1

#### Scenario: Pop from an empty stack

- **WHEN** `Pop()` is called on an empty stack
- **THEN** it SHALL return `(0, false)`

### Requirement: Peek operation

The `ArrayStack` SHALL provide a `Peek() (int, bool)` method that returns the top element without removing it.

#### Scenario: Peek at a non-empty stack

- **WHEN** `Peek()` is called on a stack with top value 5
- **THEN** it SHALL return `(5, true)` and the stack size SHALL remain unchanged

#### Scenario: Peek at an empty stack

- **WHEN** `Peek()` is called on an empty stack
- **THEN** it SHALL return `(0, false)`

### Requirement: Size operation

The `ArrayStack` SHALL provide a `Size() int` method that returns the number of elements in the stack.

#### Scenario: Size of stack

- **WHEN** `Size()` is called on a stack with 3 elements
- **THEN** it SHALL return 3

### Requirement: IsEmpty operation

The `ArrayStack` SHALL provide an `IsEmpty() bool` method that returns true when the stack has no elements.

#### Scenario: Empty stack

- **WHEN** `IsEmpty()` is called on an empty stack
- **THEN** it SHALL return true

#### Scenario: Non-empty stack

- **WHEN** `IsEmpty()` is called on a stack with elements
- **THEN** it SHALL return false

### Requirement: Print operation

The `ArrayStack` SHALL provide a `Print()` method that prints all stack elements from top to bottom to stdout.

#### Scenario: Print stack contents

- **WHEN** `Print()` is called on a stack containing [3, 2, 1] (top to bottom)
- **THEN** the elements SHALL be printed to stdout in top-to-bottom order

### Requirement: Run function displays stack operations demo

The `stack_array` package SHALL export a `Run()` function that demonstrates Push, Pop, Peek, Size, IsEmpty, and Print operations.

#### Scenario: Run is called

- **WHEN** `stack_array.Run()` is invoked
- **THEN** the function SHALL demonstrate all stack operations with printed output

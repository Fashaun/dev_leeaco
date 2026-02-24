## ADDED Requirements

### Requirement: InitArray creates an array from variadic arguments

The `array_operations` package SHALL provide an `InitArray(vals ...int) []int` function that returns a new slice containing the provided values.

#### Scenario: Initialize with values

- **WHEN** `InitArray(1, 2, 3, 4, 5)` is called
- **THEN** the function SHALL return `[1, 2, 3, 4, 5]`

#### Scenario: Initialize with no values

- **WHEN** `InitArray()` is called with no arguments
- **THEN** the function SHALL return an empty slice

### Requirement: AccessElement retrieves an element by index

The `array_operations` package SHALL provide an `AccessElement(arr []int, index int) int` function that returns the element at the given index.

#### Scenario: Access a valid index

- **WHEN** `AccessElement([]int{1, 2, 3, 4, 5}, 2)` is called
- **THEN** the function SHALL return `3`

### Requirement: InsertElement inserts a value at a given position

The `array_operations` package SHALL provide an `InsertElement(arr []int, index int, val int) []int` function that inserts `val` at `index`, shifting subsequent elements right, and returns the resulting slice.

#### Scenario: Insert at the beginning

- **WHEN** `InsertElement([]int{1, 2, 3}, 0, 0)` is called
- **THEN** the function SHALL return `[0, 1, 2, 3]`

#### Scenario: Insert in the middle

- **WHEN** `InsertElement([]int{1, 2, 4, 5}, 2, 3)` is called
- **THEN** the function SHALL return `[1, 2, 3, 4, 5]`

#### Scenario: Insert at the end

- **WHEN** `InsertElement([]int{1, 2, 3}, 3, 4)` is called
- **THEN** the function SHALL return `[1, 2, 3, 4]`

### Requirement: RemoveElement removes an element at a given position

The `array_operations` package SHALL provide a `RemoveElement(arr []int, index int) []int` function that removes the element at `index`, shifting subsequent elements left, and returns the resulting slice.

#### Scenario: Remove from the beginning

- **WHEN** `RemoveElement([]int{1, 2, 3, 4, 5}, 0)` is called
- **THEN** the function SHALL return `[2, 3, 4, 5]`

#### Scenario: Remove from the middle

- **WHEN** `RemoveElement([]int{1, 2, 3, 4, 5}, 2)` is called
- **THEN** the function SHALL return `[1, 2, 4, 5]`

#### Scenario: Remove from the end

- **WHEN** `RemoveElement([]int{1, 2, 3, 4, 5}, 4)` is called
- **THEN** the function SHALL return `[1, 2, 3, 4]`

### Requirement: TraverseArray prints each element

The `array_operations` package SHALL provide a `TraverseArray(arr []int)` function that prints each element of the array.

#### Scenario: Traverse a non-empty array

- **WHEN** `TraverseArray([]int{1, 2, 3})` is called
- **THEN** the function SHALL print each element

### Requirement: FindElement performs linear search

The `array_operations` package SHALL provide a `FindElement(arr []int, val int) int` function that returns the index of the first occurrence of `val`, or `-1` if not found.

#### Scenario: Find an existing element

- **WHEN** `FindElement([]int{1, 2, 3, 4, 5}, 3)` is called
- **THEN** the function SHALL return `2`

#### Scenario: Find a non-existing element

- **WHEN** `FindElement([]int{1, 2, 3, 4, 5}, 9)` is called
- **THEN** the function SHALL return `-1`

### Requirement: ExpandArray appends values to an array

The `array_operations` package SHALL provide an `ExpandArray(arr []int, vals ...int) []int` function that appends the given values and returns the resulting slice.

#### Scenario: Expand with multiple values

- **WHEN** `ExpandArray([]int{1, 2, 3}, 4, 5, 6)` is called
- **THEN** the function SHALL return `[1, 2, 3, 4, 5, 6]`

#### Scenario: Expand with no values

- **WHEN** `ExpandArray([]int{1, 2, 3})` is called with no additional values
- **THEN** the function SHALL return `[1, 2, 3]`

### Requirement: Run function displays array operations demo

The `array_operations` package SHALL export a `Run()` function that demonstrates all 7 array operations with printed output showing the array state after each operation.

#### Scenario: Run is called

- **WHEN** `array_operations.Run()` is invoked
- **THEN** the function SHALL print the array state after initializing, accessing, inserting, removing, traversing, finding, and expanding

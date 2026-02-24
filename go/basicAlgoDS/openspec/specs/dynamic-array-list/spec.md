### Requirement: DynamicArrayList struct with manual capacity management

The `dynamic_array_list` package SHALL provide a `DynamicArrayList` struct with `data []int`, `size int`, and `capacity int` fields, and a `NewDynamicArrayList(capacity int) *DynamicArrayList` constructor that allocates a backing array of the given capacity with size 0.

#### Scenario: Create a new dynamic array list

- **WHEN** `NewDynamicArrayList(4)` is called
- **THEN** a list with `size=0`, `capacity=4`, and an allocated backing array of length 4 SHALL be returned

### Requirement: Append adds an element to the end with automatic expansion

The `dynamic_array_list` package SHALL provide an `Append(val int)` method on `*DynamicArrayList` that adds a value to the end. When `size == capacity`, the capacity SHALL double before adding the element.

#### Scenario: Append within capacity

- **WHEN** `Append(5)` is called on a list with `size=2` and `capacity=4`
- **THEN** the element SHALL be added at index 2 and size SHALL become 3

#### Scenario: Append triggers expansion

- **WHEN** `Append(5)` is called on a list with `size=4` and `capacity=4`
- **THEN** the capacity SHALL double to 8, the element SHALL be added, and size SHALL become 5

### Requirement: Get retrieves an element by index

The `dynamic_array_list` package SHALL provide a `Get(index int) int` method on `*DynamicArrayList` that returns the element at the given index.

#### Scenario: Get a valid index

- **WHEN** `Get(1)` is called on a list containing `[10, 20, 30]`
- **THEN** the method SHALL return `20`

### Requirement: Insert adds an element at a given position

The `dynamic_array_list` package SHALL provide an `Insert(index int, val int)` method on `*DynamicArrayList` that inserts a value at the given index, shifting subsequent elements right. The capacity SHALL double if the list is full.

#### Scenario: Insert in the middle

- **WHEN** `Insert(1, 15)` is called on a list containing `[10, 20, 30]`
- **THEN** the list SHALL become `[10, 15, 20, 30]` and size SHALL increase by 1

#### Scenario: Insert at the beginning

- **WHEN** `Insert(0, 5)` is called on a list containing `[10, 20, 30]`
- **THEN** the list SHALL become `[5, 10, 20, 30]`

#### Scenario: Insert triggers expansion

- **WHEN** `Insert` is called on a full list
- **THEN** the capacity SHALL double before inserting

### Requirement: Delete removes an element at a given position

The `dynamic_array_list` package SHALL provide a `Delete(index int)` method on `*DynamicArrayList` that removes the element at the given index, shifting subsequent elements left.

#### Scenario: Delete from the middle

- **WHEN** `Delete(1)` is called on a list containing `[10, 20, 30]`
- **THEN** the list SHALL become `[10, 30]` and size SHALL decrease by 1

#### Scenario: Delete from the beginning

- **WHEN** `Delete(0)` is called on a list containing `[10, 20, 30]`
- **THEN** the list SHALL become `[20, 30]`

### Requirement: Traverse prints all elements

The `dynamic_array_list` package SHALL provide a `Traverse()` method on `*DynamicArrayList` that prints all elements in the list along with the current size and capacity.

#### Scenario: Traverse a non-empty list

- **WHEN** `Traverse()` is called on a list containing `[10, 20, 30]`
- **THEN** the method SHALL print the elements and show size and capacity

### Requirement: Concat creates a new list from two lists

The `dynamic_array_list` package SHALL provide a `Concat(other *DynamicArrayList) *DynamicArrayList` method that returns a new list containing all elements from both lists without modifying the originals.

#### Scenario: Concatenate two non-empty lists

- **WHEN** `Concat` is called with list `[1, 2, 3]` and list `[4, 5, 6]`
- **THEN** a new list `[1, 2, 3, 4, 5, 6]` SHALL be returned

#### Scenario: Original lists are unchanged

- **WHEN** `Concat` is called
- **THEN** neither the receiver list nor the argument list SHALL be modified

### Requirement: Sort arranges elements in ascending order

The `dynamic_array_list` package SHALL provide a `Sort()` method on `*DynamicArrayList` that sorts the elements in-place in ascending order using insertion sort.

#### Scenario: Sort an unsorted list

- **WHEN** `Sort()` is called on a list containing `[3, 1, 4, 1, 5]`
- **THEN** the list SHALL become `[1, 1, 3, 4, 5]`

#### Scenario: Sort an already sorted list

- **WHEN** `Sort()` is called on a list containing `[1, 2, 3]`
- **THEN** the list SHALL remain `[1, 2, 3]`

### Requirement: Run function displays dynamic array list operations demo

The `dynamic_array_list` package SHALL export a `Run()` function that demonstrates all operations including capacity expansion with printed output showing the list state and capacity after each operation.

#### Scenario: Run is called

- **WHEN** `dynamic_array_list.Run()` is invoked
- **THEN** the function SHALL print the list state after initializing, appending (with capacity changes), accessing, inserting, deleting, traversing, concatenating, and sorting

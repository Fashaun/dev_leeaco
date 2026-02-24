## Why

The project has an `array_operations` package demonstrating raw slice operations, but lacks a proper List abstraction backed by a dynamic array. A dynamic array list is a fundamental data structure that manages its own capacity, doubling when full — teaching automatic expansion, amortized O(1) appends, and the distinction between size and capacity. This complements the existing `linked_list` package by showing the array-based alternative.

## What Changes

- Add a new `dynamic_array_list` package implementing a List backed by a dynamic array with 7 core operations:
  1. Initialize a list (with optional initial values)
  2. Access elements by index
  3. Insert elements at a given position
  4. Delete elements at a given position
  5. Traverse the list (print all elements)
  6. Concatenate two lists
  7. Sort the list
- Demonstrate the expansion mechanism (capacity doubling) with printed capacity tracking
- Implement a `Run()` function demonstrating all operations
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `dynamic-array-list`: List backed by a dynamic array with initialize, access, insert, delete, traverse, concatenate, sort, and automatic expansion

### Modified Capabilities

- `cli-entry`: Add `dynamic_array_list` case to the CLI dispatcher switch statement

## Impact

- New directory: `dynamic_array_list/`
- New file: `dynamic_array_list/dynamic_array_list.go`
- Modified file: `main.go` (new import + switch case)

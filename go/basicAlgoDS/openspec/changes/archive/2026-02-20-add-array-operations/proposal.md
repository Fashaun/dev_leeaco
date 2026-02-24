## Why

The project covers sorting algorithms, graph algorithms, tree operations, and various data structures (stacks, queues, deques) but lacks a fundamental demonstration of common array operations. Arrays are the most basic data structure and a natural starting point for learners. Adding an array operations package completes the foundational coverage.

## What Changes

- Add a new `array_operations` package demonstrating 7 common array operations:
  1. Initializing arrays
  2. Accessing elements by index
  3. Inserting elements at a given position
  4. Removing elements at a given position
  5. Traversing arrays
  6. Finding elements (linear search)
  7. Expanding arrays (appending / growing capacity)
- Implement a `Run()` function that demonstrates each operation with printed output
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `array-operations`: Common array operations including initialize, access, insert, remove, traverse, find, and expand

### Modified Capabilities

- `cli-entry`: Add `array_operations` case to the CLI dispatcher switch statement

## Impact

- New directory: `array_operations/`
- New file: `array_operations/array_operations.go`
- Modified file: `main.go` (new import + switch case)

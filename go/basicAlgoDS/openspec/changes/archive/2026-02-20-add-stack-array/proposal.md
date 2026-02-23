## Why

The project already demonstrates a stack implemented via linked list. Adding an array-based (slice-based) stack implementation provides a complementary approach, showing how the same LIFO data structure can be built with different underlying storage, and helps learners compare trade-offs between array and linked-list implementations.

## What Changes

- Add a new `stack_array` package implementing a stack backed by a Go slice
- The package will provide Push, Pop, Peek, Size, IsEmpty, and Print operations
- Export a `Run()` function that demonstrates all stack operations
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `stack-array`: Stack data structure implementation using a Go slice (dynamic array), with Push, Pop, Peek, Size, IsEmpty, and Print operations

### Modified Capabilities

- `cli-entry`: Add `stack_array` as a new case in the CLI argument dispatcher

## Impact

- New directory: `stack_array/`
- New file: `stack_array/stack_array.go`
- Modified file: `main.go` (new import and switch case)

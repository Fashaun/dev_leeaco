## Why

The project has stacks (LIFO) and queues (FIFO) in both linked list and array variants, but no double-ended queue (deque). A deque allows insertion and removal from both ends, generalizing both stacks and queues. Adding an array-based deque rounds out the collection of fundamental linear data structures.

## What Changes

- Add a new `deque_array` package implementing a double-ended queue backed by a Go slice
- The package will provide PushFront, PushRear, PopFront, PopRear, PeekFront, PeekRear, Size, IsEmpty, and Print operations
- Export a `Run()` function that demonstrates all deque operations
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `deque-array`: Double-ended queue data structure implementation using a Go slice, with PushFront, PushRear, PopFront, PopRear, PeekFront, PeekRear, Size, IsEmpty, and Print operations

### Modified Capabilities

- `cli-entry`: Add `deque_array` as a new case in the CLI argument dispatcher

## Impact

- New directory: `deque_array/`
- New file: `deque_array/deque_array.go`
- Modified file: `main.go` (new import and switch case)

## Why

The project already has a queue implemented via linked list. Adding an array-based (slice-based) queue implementation provides a complementary approach, showing how the same FIFO data structure can be built with different underlying storage. This mirrors the stack pair (linked list + array) and helps learners compare trade-offs between the two approaches.

## What Changes

- Add a new `queue_array` package implementing a queue backed by a Go slice
- The package will provide Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations
- Export a `Run()` function that demonstrates all queue operations
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `queue-array`: Queue data structure implementation using a Go slice (dynamic array), with Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations

### Modified Capabilities

- `cli-entry`: Add `queue_array` as a new case in the CLI argument dispatcher

## Impact

- New directory: `queue_array/`
- New file: `queue_array/queue_array.go`
- Modified file: `main.go` (new import and switch case)

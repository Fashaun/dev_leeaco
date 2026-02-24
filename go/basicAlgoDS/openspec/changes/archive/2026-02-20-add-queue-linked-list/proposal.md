## Why

The project covers stacks (both linked list and array-based) but has no queue data structure yet. Adding a queue implemented via linked list introduces the FIFO (First-In-First-Out) pattern, complementing the existing LIFO stack examples and broadening the collection of fundamental data structures.

## What Changes

- Add a new `queue_linked_list` package implementing a queue backed by a singly linked list
- The package will provide Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations
- Export a `Run()` function that demonstrates all queue operations
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `queue-linked-list`: Queue data structure implementation using a singly linked list, with Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations

### Modified Capabilities

- `cli-entry`: Add `queue_linked_list` as a new case in the CLI argument dispatcher

## Impact

- New directory: `queue_linked_list/`
- New file: `queue_linked_list/queue_linked_list.go`
- Modified file: `main.go` (new import and switch case)

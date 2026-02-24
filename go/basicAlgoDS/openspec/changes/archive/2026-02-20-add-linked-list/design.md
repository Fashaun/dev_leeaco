## Context

The project has array-based structures (stacks, queues, deques, array operations) and tree-based structures (BST, btree traversal) but lacks a linked list. A singly linked list is the simplest pointer-based data structure and a building block for understanding more complex structures. The project already uses custom node types (e.g., `BSTNode`) so the pattern is established.

## Goals / Non-Goals

**Goals:**

- Implement a singly linked list with `ListNode` and `LinkedList` types
- Provide 5 core operations: insert, remove, access, find, and traversal/print
- Follow the established single-file package pattern with `Run()` demo function

**Non-Goals:**

- Doubly linked list (separate package if needed later)
- Generic type support (project uses `int` only)
- Circular linked list variant

## Decisions

### 1. Two types: `ListNode` and `LinkedList`

Define `ListNode` with `Val int` and `Next *ListNode`, and a `LinkedList` struct with a `Head *ListNode` pointer. Using a container struct allows methods and tracks the head cleanly.

**Alternative**: Standalone functions taking `*ListNode` like BST — works but less natural for a mutable container with a changing head pointer.

### 2. Methods on `LinkedList` rather than standalone functions

Operations like Insert, Remove, Access, Find are methods on `*LinkedList`. This is idiomatic Go for a mutable container and avoids returning updated head pointers on every call.

**Alternative**: Standalone functions returning `*ListNode` — used for BST where root changes are rare, but for linked lists the head changes frequently (insert at head, remove head).

### 3. Index-based insert and remove

`InsertAtIndex(index int, val int)` inserts before the element at `index`. `RemoveAtIndex(index int)` removes the element at `index`. Index 0 means head. This mirrors the array operations package for consistency.

### 4. Print method for traversal

`Print()` method traverses the list and prints all values in a readable format (e.g., `1 -> 2 -> 3 -> nil`). This serves as the traversal demonstration.

### 5. Demo data

Use values `[1, 2, 3, 4, 5]` inserted sequentially, then demonstrate access, find, remove, and insert operations.

## Risks / Trade-offs

- **No bounds checking**: Access and remove at invalid indices will panic. Acceptable for a learning demo — Go's convention is to panic on programmer errors.
- **No size tracking**: Could add a `Size` field but unnecessary complexity for the demo scope.

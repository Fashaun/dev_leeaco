## Why

The project has array-based and tree-based data structures but lacks a linked list implementation. A singly linked list is the most fundamental pointer-based data structure and essential for understanding node-based structures like trees and graphs. It demonstrates dynamic memory allocation, pointer manipulation, and sequential access patterns.

## What Changes

- Add a new `linked_list` package implementing a singly linked list with 5 core operations:
  1. Initializing a linked list (creating nodes)
  2. Inserting a node (at head, tail, or given index)
  3. Removing a node (by index)
  4. Accessing a node (by index)
  5. Finding a node (by value, linear search)
- Implement a `Run()` function demonstrating each operation with printed output
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `linked-list`: Singly linked list with node type, initialize, insert, remove, access, and find operations

### Modified Capabilities

- `cli-entry`: Add `linked_list` case to the CLI dispatcher switch statement

## Impact

- New directory: `linked_list/`
- New file: `linked_list/linked_list.go`
- Modified file: `main.go` (new import + switch case)

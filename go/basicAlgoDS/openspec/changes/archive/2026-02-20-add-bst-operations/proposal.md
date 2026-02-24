## Why

The project has a binary tree traversal package but lacks binary search tree (BST) operations. BST is the foundational ordered tree data structure, and its core operations — searching, inserting, and removing nodes — are essential topics in data structures education. These operations demonstrate how the BST property (left < root < right) enables efficient O(log n) average-case lookups.

## What Changes

- Add new `bst_operations` package with a BST node type and three core operations:
  - `Search(root, val)` — find a node by value
  - `Insert(root, val)` — insert a new value maintaining BST property
  - `Remove(root, val)` — remove a node handling three cases (leaf, one child, two children)
- Add in-order traversal helper to verify BST ordering
- Add `Run()` demo function that demonstrates all three operations with printed output
- Register `bst_operations` in the CLI dispatcher in `main.go`

## Capabilities

### New Capabilities

- `bst-operations`: Binary search tree operations including search, insert, remove, and in-order traversal

### Modified Capabilities

- `cli-entry`: Add `bst_operations` case to the CLI switch dispatcher

## Impact

- New file: `bst_operations/bst_operations.go`
- Modified file: `main.go` (new import and switch case)

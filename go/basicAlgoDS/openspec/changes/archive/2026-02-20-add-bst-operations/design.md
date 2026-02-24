## Context

The project has a `btree_traversal` package with a `TreeNode` struct and level-order traversal, but no BST-specific operations. BST operations (search, insert, remove) rely on the BST property: for any node, all left subtree values are less, and all right subtree values are greater. This package will be self-contained with its own node type to avoid coupling with the existing traversal package.

## Goals / Non-Goals

**Goals:**

- Implement BST search, insert, and remove operations using recursive approaches
- Include an in-order traversal helper to display the BST in sorted order
- Handle all three remove cases: leaf node, node with one child, node with two children (in-order successor)
- Match existing project conventions (single file, `Run()` demo, CLI registration)

**Non-Goals:**

- Reusing `btree_traversal.TreeNode` — keep packages decoupled
- Self-balancing trees (AVL, Red-Black) — this is a basic unbalanced BST
- Iterative implementations — recursive is clearer for educational purposes
- Duplicate value handling — assume all values are unique

## Decisions

### 1. Own node type: `BSTNode` struct

Define `BSTNode` with `Val int`, `Left *BSTNode`, `Right *BSTNode`. Separate from `btree_traversal.TreeNode` to keep packages independent and allow BST-specific semantics.

**Alternative considered:** Reuse `btree_traversal.TreeNode`. Rejected — creates coupling, and BST operations are conceptually distinct from generic tree traversal.

### 2. Recursive implementations for all operations

All three operations (search, insert, remove) use recursion. This is the textbook approach and clearly demonstrates how the BST property guides left/right decisions at each level.

### 3. Remove strategy: in-order successor for two-child case

When removing a node with two children, replace it with its in-order successor (smallest node in the right subtree). This is the standard textbook approach.

**Alternative considered:** In-order predecessor (largest in left subtree). Both are valid; successor is more commonly taught.

### 4. Function signatures

- `Search(root *BSTNode, val int) *BSTNode` — returns the found node or `nil`
- `Insert(root *BSTNode, val int) *BSTNode` — returns the (potentially new) root
- `Remove(root *BSTNode, val int) *BSTNode` — returns the (potentially new) root
- `InOrder(root *BSTNode) []int` — returns sorted values for verification

All operations return the root to handle the case where the root itself changes (e.g., removing the root node, inserting into an empty tree).

### 5. In-order traversal helper

`InOrder(root *BSTNode) []int` collects values in sorted order. This serves as both a demo output and a verification tool to confirm BST property is maintained after operations.

## Risks / Trade-offs

- [Unbalanced BST degrades to O(n)] → Acceptable for educational purposes. The demo uses a balanced-ish set of values to show typical behavior.
- [No duplicate handling] → Simplifies implementation. The demo will use unique values only.
- [Recursive approach has stack depth limit] → Not a concern for educational-sized trees.

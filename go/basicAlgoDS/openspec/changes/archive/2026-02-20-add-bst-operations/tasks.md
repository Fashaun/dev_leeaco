## 1. Create bst_operations package and node type

- [x] 1.1 Create `bst_operations/` directory and `bst_operations/bst_operations.go` with package declaration
- [x] 1.2 Define `BSTNode` struct and `NewBSTNode(val int)` constructor

## 2. Implement core BST operations

- [x] 2.1 Implement `Search(root *BSTNode, val int) *BSTNode` using recursive traversal
- [x] 2.2 Implement `Insert(root *BSTNode, val int) *BSTNode` maintaining BST property
- [x] 2.3 Implement `Remove(root *BSTNode, val int) *BSTNode` handling leaf, one-child, and two-children cases with in-order successor
- [x] 2.4 Implement `InOrder(root *BSTNode) []int` returning sorted values

## 3. Run demo function

- [x] 3.1 Implement `Run()` function demonstrating insert, search, and remove with printed BST state

## 4. Register in CLI dispatcher

- [x] 4.1 Add `bst_operations` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go bst_operations` and verify output

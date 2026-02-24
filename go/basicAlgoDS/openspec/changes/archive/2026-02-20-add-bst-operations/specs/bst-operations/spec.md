## ADDED Requirements

### Requirement: BSTNode struct

The `bst_operations` package SHALL provide a `BSTNode` struct with `Val int`, `Left *BSTNode`, and `Right *BSTNode` fields, and a `NewBSTNode(val int) *BSTNode` constructor.

#### Scenario: Create a new BST node

- **WHEN** `NewBSTNode(5)` is called
- **THEN** a node with `Val=5`, `Left=nil`, `Right=nil` SHALL be returned

### Requirement: Search function finds a node by value

The `bst_operations` package SHALL provide a `Search(root *BSTNode, val int) *BSTNode` function that returns the node with the given value, or `nil` if not found.

#### Scenario: Search for an existing value

- **WHEN** `Search` is called with a value that exists in the BST
- **THEN** the function SHALL return the node containing that value

#### Scenario: Search for a non-existing value

- **WHEN** `Search` is called with a value that does not exist in the BST
- **THEN** the function SHALL return `nil`

#### Scenario: Search in an empty tree

- **WHEN** `Search` is called with a `nil` root
- **THEN** the function SHALL return `nil`

### Requirement: Insert function adds a value maintaining BST property

The `bst_operations` package SHALL provide an `Insert(root *BSTNode, val int) *BSTNode` function that inserts a new value into the BST and returns the root. The BST property (left < root < right) SHALL be maintained.

#### Scenario: Insert into an empty tree

- **WHEN** `Insert(nil, 5)` is called
- **THEN** a new root node with `Val=5` SHALL be returned

#### Scenario: Insert a smaller value

- **WHEN** `Insert` is called with a value less than the root
- **THEN** the value SHALL be inserted into the left subtree

#### Scenario: Insert a larger value

- **WHEN** `Insert` is called with a value greater than the root
- **THEN** the value SHALL be inserted into the right subtree

#### Scenario: In-order traversal after multiple inserts

- **WHEN** values `[5, 3, 7, 1, 4]` are inserted sequentially
- **THEN** the in-order traversal SHALL produce `[1, 3, 4, 5, 7]`

### Requirement: Remove function deletes a node maintaining BST property

The `bst_operations` package SHALL provide a `Remove(root *BSTNode, val int) *BSTNode` function that removes the node with the given value and returns the root. The BST property SHALL be maintained after removal.

#### Scenario: Remove a leaf node

- **WHEN** `Remove` is called on a node with no children
- **THEN** the node SHALL be removed and its parent's pointer set to `nil`

#### Scenario: Remove a node with one child

- **WHEN** `Remove` is called on a node with exactly one child
- **THEN** the node SHALL be replaced by its child

#### Scenario: Remove a node with two children

- **WHEN** `Remove` is called on a node with two children
- **THEN** the node SHALL be replaced by its in-order successor (smallest value in the right subtree)

#### Scenario: Remove a non-existing value

- **WHEN** `Remove` is called with a value not in the BST
- **THEN** the tree SHALL remain unchanged

#### Scenario: Remove the root node

- **WHEN** `Remove` is called on the root node
- **THEN** the function SHALL return the new root after removal

### Requirement: InOrder function returns sorted values

The `bst_operations` package SHALL provide an `InOrder(root *BSTNode) []int` function that returns all values in the BST in ascending order.

#### Scenario: In-order traversal of a BST

- **WHEN** `InOrder` is called on a BST containing values `[5, 3, 7, 1, 4]`
- **THEN** the result SHALL be `[1, 3, 4, 5, 7]`

#### Scenario: In-order traversal of an empty tree

- **WHEN** `InOrder` is called with a `nil` root
- **THEN** the result SHALL be an empty slice

### Requirement: Run function displays BST operations demo

The `bst_operations` package SHALL export a `Run()` function that demonstrates search, insert, and remove operations with printed output showing the BST state after each operation.

#### Scenario: Run is called

- **WHEN** `bst_operations.Run()` is invoked
- **THEN** the function SHALL print the BST state (via in-order traversal) after inserts, search results, and the BST state after removals

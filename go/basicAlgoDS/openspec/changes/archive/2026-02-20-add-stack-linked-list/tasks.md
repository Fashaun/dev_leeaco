## 1. Stack Package Setup

- [x] 1.1 Create `stack_linked_list/` directory
- [x] 1.2 Create `stack_linked_list/stack_linked_list.go` with `package stack_linked_list`

## 2. Core Data Structure

- [x] 2.1 Define `Node` struct with `Val int` and `Next *Node` fields
- [x] 2.2 Define `LinkedListStack` struct with `top *Node` and `size int` fields

## 3. Stack Operations

- [x] 3.1 Implement `Push(val int)` method that inserts a new node at the top
- [x] 3.2 Implement `Pop() (int, bool)` method that removes and returns the top element
- [x] 3.3 Implement `Peek() (int, bool)` method that returns the top element without removing
- [x] 3.4 Implement `Size() int` method
- [x] 3.5 Implement `IsEmpty() bool` method
- [x] 3.6 Implement `Print()` method that prints elements from top to bottom

## 4. Demo Entry Point

- [x] 4.1 Implement `Run()` function that demonstrates all stack operations with printed output

## 5. CLI Registration

- [x] 5.1 Add `"basicAlgoDS/stack_linked_list"` import to `main.go`
- [x] 5.2 Add `case "stack_linked_list"` to switch in `main.go` calling `stack_linked_list.Run()`

## 6. Verification

- [x] 6.1 Run `go run main.go stack_linked_list` and verify stack operations demo output
- [x] 6.2 Run `go run main.go unknown` and verify error message still works

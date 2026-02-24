## 1. Create queue_linked_list package

- [x] 1.1 Create `queue_linked_list/` directory and `queue_linked_list/queue_linked_list.go` with package declaration
- [x] 1.2 Define `Node` struct with `Val int` and `Next *Node` fields
- [x] 1.3 Define `LinkedListQueue` struct with `front *Node`, `rear *Node`, and `size int` fields
- [x] 1.4 Implement `Enqueue(val int)` method that inserts at rear
- [x] 1.5 Implement `Dequeue() (int, bool)` method that removes from front with empty-queue guard
- [x] 1.6 Implement `Peek() (int, bool)` method with empty-queue guard
- [x] 1.7 Implement `Size() int` and `IsEmpty() bool` methods
- [x] 1.8 Implement `Print()` method that prints elements from front to rear

## 2. Run demo function

- [x] 2.1 Implement `Run()` function demonstrating Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations

## 3. Register in CLI dispatcher

- [x] 3.1 Add `queue_linked_list` import and switch case in `main.go`

## 4. Verify

- [x] 4.1 Run `go run main.go queue_linked_list` and verify output

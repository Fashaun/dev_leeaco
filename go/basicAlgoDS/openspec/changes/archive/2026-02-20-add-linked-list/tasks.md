## 1. Create linked_list package and types

- [x] 1.1 Create `linked_list/` directory and `linked_list/linked_list.go` with package declaration
- [x] 1.2 Define `ListNode` struct, `NewListNode(val int)` constructor, `LinkedList` struct, and `NewLinkedList()` constructor

## 2. Implement core operations

- [x] 2.1 Implement `InsertAtIndex(index int, val int)` method handling head, middle, and tail insertion
- [x] 2.2 Implement `RemoveAtIndex(index int)` method handling head, middle, and tail removal
- [x] 2.3 Implement `Access(index int) int` method returning the value at the given index
- [x] 2.4 Implement `Find(val int) int` method returning index of first match or -1
- [x] 2.5 Implement `Print()` method displaying list as `val1 -> val2 -> ... -> nil`

## 3. Implement demo function

- [x] 3.1 Implement `Run()` function demonstrating all operations with printed output

## 4. Register in CLI dispatcher

- [x] 4.1 Add `linked_list` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go linked_list` and verify output shows all operations

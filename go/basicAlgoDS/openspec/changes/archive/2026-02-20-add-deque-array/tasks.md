## 1. Create deque_array package

- [x] 1.1 Create `deque_array/` directory and `deque_array/deque_array.go` with package declaration
- [x] 1.2 Define `ArrayDeque` struct with `data []int` slice and `size int` field
- [x] 1.3 Implement `PushFront(val int)` method using prepend
- [x] 1.4 Implement `PushRear(val int)` method using `append()`
- [x] 1.5 Implement `PopFront() (int, bool)` method with reslicing and empty-deque guard
- [x] 1.6 Implement `PopRear() (int, bool)` method with reslicing and empty-deque guard
- [x] 1.7 Implement `PeekFront() (int, bool)` method with empty-deque guard
- [x] 1.8 Implement `PeekRear() (int, bool)` method with empty-deque guard
- [x] 1.9 Implement `Size() int` and `IsEmpty() bool` methods
- [x] 1.10 Implement `Print()` method that prints elements from front to rear

## 2. Run demo function

- [x] 2.1 Implement `Run()` function demonstrating PushFront, PushRear, PopFront, PopRear, PeekFront, PeekRear, Size, IsEmpty, and Print operations

## 3. Register in CLI dispatcher

- [x] 3.1 Add `deque_array` import and switch case in `main.go`

## 4. Verify

- [x] 4.1 Run `go run main.go deque_array` and verify output

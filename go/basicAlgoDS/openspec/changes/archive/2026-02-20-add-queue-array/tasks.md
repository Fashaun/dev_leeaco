## 1. Create queue_array package

- [x] 1.1 Create `queue_array/` directory and `queue_array/queue_array.go` with package declaration
- [x] 1.2 Define `ArrayQueue` struct with `data []int` slice and `size int` field
- [x] 1.3 Implement `Enqueue(val int)` method using `append()`
- [x] 1.4 Implement `Dequeue() (int, bool)` method with reslicing and empty-queue guard
- [x] 1.5 Implement `Peek() (int, bool)` method with empty-queue guard
- [x] 1.6 Implement `Size() int` and `IsEmpty() bool` methods
- [x] 1.7 Implement `Print()` method that prints elements from front to rear

## 2. Run demo function

- [x] 2.1 Implement `Run()` function demonstrating Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations

## 3. Register in CLI dispatcher

- [x] 3.1 Add `queue_array` import and switch case in `main.go`

## 4. Verify

- [x] 4.1 Run `go run main.go queue_array` and verify output

## 1. Create stack_array package

- [x] 1.1 Create `stack_array/` directory and `stack_array/stack_array.go` with package declaration
- [x] 1.2 Define `ArrayStack` struct with `data []int` slice and `size int` field
- [x] 1.3 Implement `Push(val int)` method using `append()`
- [x] 1.4 Implement `Pop() (int, bool)` method with reslicing and empty-stack guard
- [x] 1.5 Implement `Peek() (int, bool)` method with empty-stack guard
- [x] 1.6 Implement `Size() int` and `IsEmpty() bool` methods
- [x] 1.7 Implement `Print()` method that prints elements from top to bottom

## 2. Run demo function

- [x] 2.1 Implement `Run()` function demonstrating Push, Pop, Peek, Size, IsEmpty, and Print operations

## 3. Register in CLI dispatcher

- [x] 3.1 Add `stack_array` import and switch case in `main.go`

## 4. Verify

- [x] 4.1 Run `go run main.go stack_array` and verify output

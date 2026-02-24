## 1. Create dynamic_array_list package and struct

- [x] 1.1 Create `dynamic_array_list/` directory and `dynamic_array_list/dynamic_array_list.go` with package declaration
- [x] 1.2 Define `DynamicArrayList` struct with `data []int`, `size int`, `capacity int` fields and `NewDynamicArrayList(capacity int)` constructor

## 2. Implement expansion and core operations

- [x] 2.1 Implement unexported `expand()` method that doubles capacity and copies elements
- [x] 2.2 Implement `Append(val int)` method with automatic expansion when full
- [x] 2.3 Implement `Get(index int) int` method returning element at index
- [x] 2.4 Implement `Insert(index int, val int)` method shifting elements right with expansion
- [x] 2.5 Implement `Delete(index int)` method shifting elements left

## 3. Implement traversal, concatenation, and sort

- [x] 3.1 Implement `Traverse()` method printing elements with size and capacity
- [x] 3.2 Implement `Concat(other *DynamicArrayList) *DynamicArrayList` returning a new combined list
- [x] 3.3 Implement `Sort()` method using in-place insertion sort

## 4. Implement demo function

- [x] 4.1 Implement `Run()` function demonstrating all operations with capacity tracking output

## 5. Register in CLI dispatcher

- [x] 5.1 Add `dynamic_array_list` import and switch case in `main.go`

## 6. Verify

- [x] 6.1 Run `go run main.go dynamic_array_list` and verify output shows all operations including capacity expansion

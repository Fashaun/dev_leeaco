## 1. Create array_operations package and file

- [x] 1.1 Create `array_operations/` directory and `array_operations/array_operations.go` with package declaration

## 2. Implement array operation functions

- [x] 2.1 Implement `InitArray(vals ...int) []int` returning a new slice from variadic args
- [x] 2.2 Implement `AccessElement(arr []int, index int) int` returning element at index
- [x] 2.3 Implement `InsertElement(arr []int, index int, val int) []int` inserting at position and returning new slice
- [x] 2.4 Implement `RemoveElement(arr []int, index int) []int` removing at position and returning new slice
- [x] 2.5 Implement `TraverseArray(arr []int)` printing each element
- [x] 2.6 Implement `FindElement(arr []int, val int) int` returning index or -1
- [x] 2.7 Implement `ExpandArray(arr []int, vals ...int) []int` appending values and returning new slice

## 3. Implement demo function

- [x] 3.1 Implement `Run()` function demonstrating all 7 operations with printed output

## 4. Register in CLI dispatcher

- [x] 4.1 Add `array_operations` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go array_operations` and verify output shows all operations

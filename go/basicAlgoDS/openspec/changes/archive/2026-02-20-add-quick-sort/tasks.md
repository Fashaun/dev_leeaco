## 1. Create quick_sort package and partition function

- [x] 1.1 Create `quick_sort/` directory and `quick_sort/quick_sort.go` with package declaration
- [x] 1.2 Implement unexported `partition(arr []int, low, high int) int` using Lomuto scheme (last element as pivot)

## 2. Implement recursive sort and exported function

- [x] 2.1 Implement unexported `quickSort(arr []int, low, high int)` that partitions and recurses on left/right subarrays
- [x] 2.2 Implement exported `QuickSort(arr []int)` that delegates to `quickSort(arr, 0, len(arr)-1)` with empty-slice guard

## 3. Implement demo function

- [x] 3.1 Implement `Run()` function that prints before/after sorting of `[5, 3, 8, 1, 2]`

## 4. Register in CLI dispatcher

- [x] 4.1 Add `quick_sort` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go quick_sort` and verify output shows `[5 3 8 1 2]` → `[1 2 3 5 8]`

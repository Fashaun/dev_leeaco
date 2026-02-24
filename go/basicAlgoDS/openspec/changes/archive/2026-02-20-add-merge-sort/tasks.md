## 1. Create merge_sort package and file

- [x] 1.1 Create `merge_sort/` directory and `merge_sort/merge_sort.go` with package declaration

## 2. Implement core functions

- [x] 2.1 Implement unexported `merge(left, right []int) []int` that merges two sorted slices into one
- [x] 2.2 Implement `MergeSort(arr []int) []int` using recursive divide-and-conquer that returns a new sorted slice

## 3. Implement demo function

- [x] 3.1 Implement `Run()` function printing before/after sorting with demo data `[5, 3, 8, 1, 2]`

## 4. Register in CLI dispatcher

- [x] 4.1 Add `merge_sort` import and switch case in `main.go`

## 5. Verify

- [x] 5.1 Run `go run main.go merge_sort` and verify output shows `[1, 2, 3, 5, 8]`

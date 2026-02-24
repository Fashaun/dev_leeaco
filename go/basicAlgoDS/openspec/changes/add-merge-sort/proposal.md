## Why

The project includes O(n^2) sorting algorithms (bubble sort, selection sort, insertion sort) but lacks any O(n log n) sorting algorithm. Merge sort is a fundamental divide-and-conquer algorithm that guarantees O(n log n) time complexity in all cases (best, average, worst), making it an essential addition to the algorithm collection.

## What Changes

- Add a new `merge_sort` package implementing the merge sort algorithm
- Implement `MergeSort(arr []int) []int` using recursive divide-and-conquer with a `merge` helper
- Implement `Run()` function demonstrating the sort with printed before/after output
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `merge-sort`: Merge sort algorithm with recursive split and merge, returning a new sorted slice

### Modified Capabilities

- `cli-entry`: Add `merge_sort` case to the CLI dispatcher switch statement

## Impact

- New directory: `merge_sort/`
- New file: `merge_sort/merge_sort.go`
- Modified file: `main.go` (new import + switch case)

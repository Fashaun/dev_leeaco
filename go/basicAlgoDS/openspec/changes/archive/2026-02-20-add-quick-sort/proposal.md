## Why

The project has bubble sort, selection sort, insertion sort (all O(n²)), and merge sort (O(n log n) but allocates new slices). Quick sort fills a critical gap as the most widely-used in-place divide-and-conquer sorting algorithm with O(n log n) average performance, completing the collection of fundamental comparison sorts.

## What Changes

- Add a new `quick_sort` package with an in-place quick sort implementation using the Lomuto partition scheme
- Export `QuickSort(arr []int)` that sorts in-place and a `Run()` demo function
- Register the package in the CLI dispatcher (`main.go`)

## Capabilities

### New Capabilities

- `quick-sort`: In-place quick sort with Lomuto partitioning, recursive divide-and-conquer, and demo output

### Modified Capabilities

- `cli-entry`: Add `quick_sort` case to the CLI dispatcher switch

## Impact

- New directory: `quick_sort/`
- New file: `quick_sort/quick_sort.go`
- Modified file: `main.go` (import + switch case)

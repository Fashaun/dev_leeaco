## Why

The project has bubble sort and selection sort but is missing insertion sort — the third fundamental O(n^2) comparison sort. Insertion sort is notable for being adaptive (O(n) on nearly sorted data) and stable, making it a practical choice for small or partially sorted datasets.

## What Changes

- Add new `insertion_sort` package with `InsertionSort(arr []int)` function that sorts in-place in ascending order
- Add `Run()` demo function that prints array before and after sorting
- Register `insertion_sort` in the CLI dispatcher in `main.go`

## Capabilities

### New Capabilities

- `insertion-sort`: Insertion sort algorithm implementation with in-place ascending sort and demo output

### Modified Capabilities

- `cli-entry`: Add `insertion_sort` case to the CLI switch dispatcher

## Impact

- New file: `insertion_sort/insertion_sort.go`
- Modified file: `main.go` (new import and switch case)

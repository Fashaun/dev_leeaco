## Why

The project has bubble sort but no other comparison-based sorting algorithms. Adding selection sort introduces another fundamental O(n^2) sorting algorithm that uses a different strategy — finding the minimum element and placing it in position — complementing bubble sort's adjacent-swap approach and helping learners compare sorting techniques.

## What Changes

- Add a new `selection_sort` package implementing the selection sort algorithm
- The package will provide a `SelectionSort(arr []int)` function that sorts in-place in ascending order
- Export a `Run()` function that demonstrates the sorting with before/after output
- Register the new package in `main.go` CLI dispatcher

## Capabilities

### New Capabilities

- `selection-sort`: Selection sort algorithm implementation that sorts an `[]int` slice in-place in ascending order, with a `Run()` demo function

### Modified Capabilities

- `cli-entry`: Add `selection_sort` as a new case in the CLI argument dispatcher

## Impact

- New directory: `selection_sort/`
- New file: `selection_sort/selection_sort.go`
- Modified file: `main.go` (new import and switch case)

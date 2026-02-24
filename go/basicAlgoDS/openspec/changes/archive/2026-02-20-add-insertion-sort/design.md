## Context

The project already has two O(n^2) sorting algorithms (bubble sort and selection sort), each as a standalone package with an exported sort function and `Run()` demo. Insertion sort follows the same pattern — a single `.go` file in its own package directory, registered in `main.go`.

## Goals / Non-Goals

**Goals:**

- Implement insertion sort as an in-place, ascending-order sort on `[]int`
- Match the existing sorting package conventions (function signature, `Run()` demo, CLI registration)

**Non-Goals:**

- Generic/type-parameterized sorting — project uses `int` only
- Descending order or custom comparators
- Performance benchmarks or optimizations (binary insertion sort, etc.)

## Decisions

### 1. Algorithm: standard shifting insertion sort

Use the classic approach: iterate from index 1, save the current element as `key`, shift larger elements right, insert `key` at the correct position. This is the textbook insertion sort and clearly demonstrates the algorithm.

**Alternative considered:** Binary insertion sort (use binary search to find insertion point). Rejected — adds complexity without meaningful benefit for an educational example.

### 2. Package structure: single file `insertion_sort/insertion_sort.go`

Matches `bubble_sort` and `selection_sort` patterns. Function name `InsertionSort(arr []int)` with in-place modification.

### 3. Demo data: `[]int{5, 3, 8, 1, 2}`

Same test array as selection sort for consistency across sorting demos.

## Risks / Trade-offs

- [Shifting is O(n) per insertion] → Acceptable for educational purposes; insertion sort's O(n^2) worst case is well-known and expected.
- [No stability test] → Insertion sort is stable by nature (equal elements maintain order), but since we sort `int` there's no observable difference. Not testing this keeps things simple.

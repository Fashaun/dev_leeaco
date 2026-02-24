## Context

The project has four sorting algorithms: bubble sort, selection sort, and insertion sort (all O(n²) in-place), and merge sort (O(n log n) but returns a new slice). Quick sort adds the first in-place O(n log n) average-case sort, using the classic divide-and-conquer partitioning approach.

## Goals / Non-Goals

**Goals:**

- Implement a clear, educational quick sort with visible partitioning logic
- Sort in-place (mutate the input slice) consistent with bubble/selection/insertion sort
- Use the same demo data `[5, 3, 8, 1, 2]` and before/after output pattern

**Non-Goals:**

- Optimized pivot selection (median-of-three, randomized) — Lomuto with last-element pivot is sufficient for educational purposes
- Handling duplicate-heavy inputs efficiently (three-way partitioning)
- Generic type support — `int` only per project convention

## Decisions

**1. Lomuto partition scheme (last element as pivot)**

Lomuto is simpler to understand and implement than Hoare's partition. The pivot is always the last element of the current subarray. Elements are rearranged so that all values ≤ pivot come before it and all values > pivot come after. The partition function returns the final pivot index.

Alternative: Hoare's partition (two pointers from both ends) is more efficient in practice but harder to follow for educational purposes.

**2. In-place sorting via `QuickSort(arr []int)`**

The exported function sorts the slice in-place, matching the pattern of bubble sort, selection sort, and insertion sort. An unexported `quickSort(arr []int, low, high int)` handles the recursive subdivision, and an unexported `partition(arr []int, low, high int) int` performs the Lomuto partitioning.

Alternative: Returning a new sorted slice (like merge sort) — rejected because quick sort is inherently in-place and the project already has merge sort for the "returns new slice" pattern.

**3. Single file structure**

All code in `quick_sort/quick_sort.go`: exported `QuickSort`, unexported `quickSort` and `partition` helpers, and exported `Run()`. Matches existing sort packages.

## Risks / Trade-offs

- [O(n²) worst case with sorted input] → Acceptable for educational purposes; Lomuto with last-element pivot degrades on already-sorted data, but the demo data `[5, 3, 8, 1, 2]` avoids this.
- [Not stable sort] → Expected behavior for quick sort; the project's merge sort already provides a stable option.

## Context

The project has bubble sort as its only sorting algorithm. Selection sort is another fundamental O(n^2) comparison sort that works differently: it repeatedly finds the minimum element from the unsorted portion and swaps it into position. This provides a useful contrast to bubble sort's adjacent-swap strategy.

## Goals / Non-Goals

**Goals:**

- Provide an in-place selection sort implementation on `[]int`
- Follow the same pattern as `bubble_sort` (exported sort function + `Run()` demo)
- Register in `main.go` CLI dispatcher

**Non-Goals:**

- Generic/type-parameterized sorting (keep it `[]int` only, matching `bubble_sort`)
- Optimized variants (e.g., double selection sort selecting both min and max)
- Stability guarantees (standard selection sort is not stable)

## Decisions

**1. In-place sorting with swap**

For each position `i` from `0` to `n-2`, find the index of the minimum element in `arr[i:]` and swap it with `arr[i]`. This matches the classic textbook selection sort. Alternative: copy-based approach — rejected because in-place is simpler and consistent with `bubble_sort`.

**2. Function signature: `SelectionSort(arr []int)`**

Matches `BubbleSort(arr []int)` in the `bubble_sort` package. Sorts the slice in-place in ascending order.

**3. Run() demo mirrors bubble_sort.Run()**

Print the array before and after sorting to show the algorithm's effect, keeping the demo style consistent across sorting algorithm packages.

## Risks / Trade-offs

- [O(n^2) time complexity] → Expected for an educational demo of a simple sorting algorithm.
- [Not stable] → Acceptable; stability is not a concern for this educational example.

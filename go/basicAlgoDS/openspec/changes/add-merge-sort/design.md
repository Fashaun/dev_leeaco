## Context

The project has three O(n^2) sorting algorithms (bubble, selection, insertion sort). Merge sort is the first O(n log n) algorithm to be added. All existing sort packages follow the same pattern: single `.go` file, exported sort function operating on `[]int`, and a `Run()` demo function using the same sample data `[5, 3, 8, 1, 2]`.

## Goals / Non-Goals

**Goals:**

- Implement a correct, readable recursive merge sort
- Follow the established package conventions (single file, `Run()` function, same demo data)
- Register in CLI dispatcher

**Non-Goals:**

- In-place merge sort (standard merge sort uses O(n) auxiliary space)
- Iterative (bottom-up) merge sort
- Generic type support (project uses `int` only)

## Decisions

### 1. Return a new sorted slice rather than sort in-place

Merge sort naturally produces new slices during the merge step. Rather than fighting this with index manipulation, `MergeSort` will return `[]int`. The `Run()` function will show both before and after. This differs from the in-place pattern of bubble/selection/insertion sort but is idiomatic for merge sort.

**Alternative**: Sort in-place using indices — adds complexity without clarity benefit for a learning project.

### 2. Top-down recursive approach

Use the standard recursive divide-and-conquer: split array in half, recurse on each half, merge the two sorted halves. This is the most intuitive and commonly taught version.

**Alternative**: Bottom-up iterative — more efficient in practice but harder to understand.

### 3. Separate `merge` helper as unexported function

Extract the merge step into an unexported `merge(left, right []int) []int` function to keep `MergeSort` clean and focused on the divide step.

### 4. Same demo data as other sorts

Use `[5, 3, 8, 1, 2]` → `[1, 2, 3, 5, 8]` for consistency across all sorting demos.

## Risks / Trade-offs

- **O(n) extra space**: Merge sort allocates new slices during merge. Acceptable for a learning project focused on correctness and clarity.
- **Return vs in-place inconsistency**: `MergeSort` returns a new slice while other sorts modify in-place. The `Run()` output will make this clear.

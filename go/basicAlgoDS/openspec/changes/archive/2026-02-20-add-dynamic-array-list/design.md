## Context

The project has `array_operations` (raw slice functions) and `linked_list` (pointer-based list). This package adds a proper List abstraction backed by a dynamic array, managing its own `data []int`, `size`, and `capacity` fields explicitly. Unlike `array_operations` which uses Go slices directly, this demonstrates the classic dynamic array pattern: manual capacity management with doubling on overflow.

## Goals / Non-Goals

**Goals:**

- Implement a `DynamicArrayList` struct with explicit `size` and `capacity` tracking
- Provide 7 operations: initialize, access, insert, delete, traverse, concatenate, sort
- Demonstrate capacity expansion (doubling) with printed capacity changes
- Follow the single-file package pattern with `Run()` demo function

**Non-Goals:**

- Shrinking capacity when elements are removed (keep it simple)
- Generic type support (project uses `int` only)
- Thread safety

## Decisions

### 1. Custom struct with manual capacity management

Define `DynamicArrayList` with `data []int`, `size int`, and `capacity int` fields. Do NOT rely on Go's built-in slice growth — allocate a fixed-size backing array and manage growth manually. This is the educational point of this data structure.

**Alternative**: Wrap Go slices with `append` — defeats the purpose of demonstrating dynamic array mechanics.

### 2. Double capacity on overflow

When `size == capacity`, allocate a new array with `2 * capacity` and copy elements. Start with initial capacity of 4. Print capacity changes so the expansion is visible in demo output.

### 3. Methods on `*DynamicArrayList`

All operations are methods: `Get(index)`, `Insert(index, val)`, `Delete(index)`, `Traverse()`, `Concat(other)`, `Sort()`, `Append(val)`. Use a container struct since size/capacity change frequently.

### 4. Sort uses insertion sort

Reuse the insertion sort algorithm (already in the project) by implementing it inline. Simple, in-place, and consistent with the project's educational focus.

**Alternative**: Call `sort.Ints` — hides the sorting logic, less educational.

### 5. Concat creates a new list

`Concat(other *DynamicArrayList) *DynamicArrayList` returns a new list containing elements from both lists. Does not modify either original.

### 6. Demo data

Initialize with `[1, 2, 3]`, then demonstrate append (showing capacity growth), insert, delete, access, traverse, concat, and sort.

## Risks / Trade-offs

- **Manual memory management overhead**: More code than using Go slices, but that's the educational value.
- **No shrink**: Capacity never decreases. Acceptable for a learning demo.
- **Inline sort**: Duplicates insertion sort logic rather than importing from `insertion_sort` package. Keeps the package self-contained with no cross-package dependencies.

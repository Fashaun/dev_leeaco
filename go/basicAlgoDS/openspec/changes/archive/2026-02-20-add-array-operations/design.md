## Context

The project demonstrates data structures and algorithms using Go. Arrays (slices in Go) are the most fundamental data structure but have no dedicated package yet. This package focuses on demonstrating common operations rather than implementing a custom array type — Go's built-in slices already provide the underlying mechanics.

## Goals / Non-Goals

**Goals:**

- Demonstrate 7 common array operations using Go slices: initialize, access, insert, remove, traverse, find, expand
- Each operation is an exported function operating on `[]int`
- Follow the established single-file package pattern with `Run()` demo function

**Non-Goals:**

- Custom array struct wrapping a slice (Go slices are sufficient)
- Generic type support (project uses `int` only)
- Bounds-checking wrappers (Go already panics on out-of-bounds access)

## Decisions

### 1. Use `[]int` directly, no wrapper type

Go slices already handle dynamic sizing, capacity management, and bounds checking. Wrapping them adds no educational value. Each function takes and/or returns `[]int`.

**Alternative**: Custom `Array` struct — adds unnecessary indirection for a demo package.

### 2. Each operation as a standalone exported function

Provide individual functions for each operation so each can be demonstrated independently in `Run()`:
- `InitArray(vals ...int) []int` — variadic constructor
- `AccessElement(arr []int, index int) int` — access by index
- `InsertElement(arr []int, index int, val int) []int` — insert at position, shifting elements right
- `RemoveElement(arr []int, index int) []int` — remove at position, shifting elements left
- `TraverseArray(arr []int)` — print each element
- `FindElement(arr []int, val int) int` — linear search returning index or -1
- `ExpandArray(arr []int, vals ...int) []int` — append values

**Alternative**: Methods on a struct — not idiomatic for a demo of raw array/slice operations.

### 3. Insert and Remove return new slices

`InsertElement` and `RemoveElement` return the modified slice since Go's `append` may allocate a new underlying array. This is consistent with Go slice semantics.

### 4. Same demo data pattern

Use a simple initial array like `[1, 2, 3, 4, 5]` and demonstrate each operation sequentially with printed output.

## Risks / Trade-offs

- **Simplicity vs completeness**: Functions are intentionally simple (no error returns for out-of-bounds). Go's runtime panic is sufficient for a learning demo.
- **Overlap with Go built-ins**: Operations like `append` and indexing are built-in. The value is in showing them as named, explicit operations for learners.

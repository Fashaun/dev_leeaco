## Context

The project has an existing `stack_linked_list` package that implements a LIFO stack using a singly linked list. This change adds a parallel implementation using a Go slice as the backing store, following the same API pattern (Push, Pop, Peek, Size, IsEmpty, Print, Run).

## Goals / Non-Goals

**Goals:**

- Provide an array-based stack implementation with the same operations as `stack_linked_list`
- Use a Go slice (`[]int`) as the underlying storage
- Follow the project's established conventions (package name = directory name, exported `Run()` function)
- Register in `main.go` CLI dispatcher

**Non-Goals:**

- Generic/type-parameterized stack (keep it `int` only, matching existing style)
- Interface abstraction over both stack implementations
- Benchmarking or performance comparison tooling

## Decisions

**1. Use Go slice (`[]int`) as backing store**

The slice grows dynamically via `append()`, with the top of the stack at the end of the slice. This gives O(1) amortized Push and O(1) Pop (by reslicing). Alternative: fixed-size array with manual capacity management — rejected because Go slices already handle dynamic resizing and this is an educational project.

**2. Match the `stack_linked_list` API surface**

Use the same method signatures: `Push(val int)`, `Pop() (int, bool)`, `Peek() (int, bool)`, `Size() int`, `IsEmpty() bool`, `Print()`. This makes it easy for learners to compare the two implementations side by side.

**3. Struct name: `ArrayStack`**

Name the struct `ArrayStack` to parallel `LinkedListStack` in the other package.

**4. Pop/Peek return `(int, bool)` comma-ok pattern**

Consistent with `stack_linked_list`. Returns `(0, false)` on empty stack instead of panicking.

## Risks / Trade-offs

- [Slice memory not shrinking on Pop] → Acceptable for an educational demo; the slice capacity remains after popping but this is fine for small demos.
- [No thread safety] → Consistent with the linked list implementation; out of scope for this project.

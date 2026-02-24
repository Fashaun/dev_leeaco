## Context

The project has a queue implemented via linked list (`queue_linked_list`) and stacks in both linked list and array variants. This change adds an array-based queue to complete the pair, following the same project conventions.

## Goals / Non-Goals

**Goals:**

- Provide an array-based queue with the same operations as `queue_linked_list`
- Use a Go slice (`[]int`) as the underlying storage
- Follow the project's established conventions (package name = directory name, exported `Run()` function)
- Register in `main.go` CLI dispatcher

**Non-Goals:**

- Circular buffer / ring buffer implementation (keep it simple with slice append and reslice)
- Generic/type-parameterized queue (keep it `int` only, matching existing style)
- Interface abstraction over both queue implementations

## Decisions

**1. Use Go slice (`[]int`) as backing store**

Enqueue appends to the end of the slice via `append()`. Dequeue returns the first element and reslices `data[1:]`. This is O(1) amortized for Enqueue and O(n) for Dequeue due to element shifting. Alternative: circular buffer — rejected because this is an educational demo prioritizing simplicity over performance.

**2. Struct name: `ArrayQueue`**

Parallels `ArrayStack` and `LinkedListQueue` in naming convention.

**3. Match the `queue_linked_list` API surface**

Use the same method signatures: `Enqueue(val int)`, `Dequeue() (int, bool)`, `Peek() (int, bool)`, `Size() int`, `IsEmpty() bool`, `Print()`. This makes it easy for learners to compare the two implementations.

**4. Dequeue/Peek return `(int, bool)` comma-ok pattern**

Consistent with all other data structure implementations in the project. Returns `(0, false)` on empty queue.

## Risks / Trade-offs

- [Dequeue is O(n) due to slice reslicing] → Acceptable for an educational demo; a circular buffer would be O(1) but adds complexity.
- [Slice memory not shrinking] → After dequeue, the underlying array capacity remains; acceptable for small demos.
- [No thread safety] → Consistent with all other implementations; out of scope.

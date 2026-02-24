## Context

The project has stacks and queues in both linked list and array variants. A deque (double-ended queue) generalizes both by allowing insertion and removal at both ends. This change adds an array-based deque using a Go slice, following the same project conventions.

## Goals / Non-Goals

**Goals:**

- Provide an array-based deque with PushFront, PushRear, PopFront, PopRear, PeekFront, PeekRear, Size, IsEmpty, and Print operations
- Use a Go slice (`[]int`) as the underlying storage
- Follow the project's established conventions (package name = directory name, exported `Run()` function)
- Register in `main.go` CLI dispatcher

**Non-Goals:**

- Circular buffer implementation (keep it simple with slice operations)
- Generic/type-parameterized deque (keep it `int` only, matching existing style)
- Interface abstraction over deque implementations

## Decisions

**1. Use Go slice (`[]int`) as backing store**

PushRear appends via `append()`. PushFront prepends using `append([]int{val}, data...)`. PopFront reslices `data[1:]`. PopRear reslices `data[:len-1]`. This prioritizes simplicity over performance. Alternative: circular buffer with head/tail indices — rejected because it adds complexity not needed for an educational demo.

**2. Struct name: `ArrayDeque`**

Follows the naming convention: `ArrayStack`, `ArrayQueue`, `ArrayDeque`.

**3. All Pop/Peek methods return `(int, bool)` comma-ok pattern**

Consistent with all other data structure implementations. Returns `(0, false)` on empty deque.

**4. Print order: front to rear**

Print elements left-to-right from front to rear, matching the natural slice order after operations.

## Risks / Trade-offs

- [PushFront is O(n) due to prepend] → Acceptable for an educational demo; a circular buffer would be O(1) amortized but adds complexity.
- [PopFront is O(n) due to reslicing] → Same trade-off as `queue_array`; acceptable here.
- [No thread safety] → Consistent with all other implementations; out of scope.

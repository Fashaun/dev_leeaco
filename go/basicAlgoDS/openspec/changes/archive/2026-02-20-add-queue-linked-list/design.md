## Context

The project has stack implementations (linked list and array) but no queue data structure. A queue is a fundamental FIFO data structure that complements the existing LIFO stacks. This change adds a linked-list-backed queue following the same project conventions.

## Goals / Non-Goals

**Goals:**

- Provide a linked-list-based queue with Enqueue, Dequeue, Peek, Size, IsEmpty, and Print operations
- Use a singly linked list with both `front` and `rear` pointers for O(1) enqueue and dequeue
- Follow the project's established conventions (package name = directory name, exported `Run()` function)
- Register in `main.go` CLI dispatcher

**Non-Goals:**

- Generic/type-parameterized queue (keep it `int` only, matching existing style)
- Circular queue or double-ended queue (deque)
- Interface abstraction over queue implementations

## Decisions

**1. Singly linked list with `front` and `rear` pointers**

Enqueue at the rear, dequeue from the front. Both operations are O(1). Alternative: use only a `front` pointer and traverse to the tail on enqueue — rejected because that makes Enqueue O(n).

**2. Struct name: `LinkedListQueue`**

Parallels `LinkedListStack` in naming convention. The struct holds `front *Node`, `rear *Node`, and `size int`.

**3. Reuse `Node` struct pattern**

Define a `Node` struct with `Val int` and `Next *Node`, consistent with the `stack_linked_list` package. Each package defines its own `Node` since Go packages are self-contained.

**4. Dequeue/Peek return `(int, bool)` comma-ok pattern**

Consistent with the stack implementations. Returns `(0, false)` on empty queue instead of panicking.

**5. Print order: front to rear**

Print elements in FIFO order (front first, rear last), which is the natural traversal order of the linked list.

## Risks / Trade-offs

- [No thread safety] → Consistent with all other implementations in the project; out of scope.
- [Memory not pooled] → Each Enqueue allocates a new Node; acceptable for an educational demo.

package queue_array

import "fmt"

type ArrayQueue struct {
	data []int
	size int
}

func (q *ArrayQueue) Enqueue(val int) {
	q.data = append(q.data, val)
	q.size++
}

func (q *ArrayQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	q.size--
	return val, true
}

func (q *ArrayQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.data[0], true
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) Print() {
	fmt.Print("Queue (front -> rear): [")
	for i := 0; i < q.size; i++ {
		fmt.Print(q.data[i])
		if i < q.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func Run() {
	q := &ArrayQueue{}

	fmt.Println("=== Enqueue 1, 2, 3, 4, 5 ===")
	for _, v := range []int{1, 2, 3, 4, 5} {
		q.Enqueue(v)
	}
	q.Print()
	fmt.Printf("Size: %d\n", q.Size())
	fmt.Printf("IsEmpty: %v\n", q.IsEmpty())

	if val, ok := q.Peek(); ok {
		fmt.Printf("\n=== Peek: %d ===\n", val)
	}
	q.Print()

	fmt.Println("\n=== Dequeue ===")
	if val, ok := q.Dequeue(); ok {
		fmt.Printf("Dequeued: %d\n", val)
	}
	q.Print()
	fmt.Printf("Size: %d\n", q.Size())

	fmt.Println("\n=== Dequeue all remaining ===")
	for !q.IsEmpty() {
		if val, ok := q.Dequeue(); ok {
			fmt.Printf("Dequeued: %d\n", val)
		}
	}
	q.Print()
	fmt.Printf("IsEmpty: %v\n", q.IsEmpty())

	fmt.Println("\n=== Dequeue from empty queue ===")
	_, ok := q.Dequeue()
	fmt.Printf("Dequeue success: %v\n", ok)
}

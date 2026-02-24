package queue_linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *LinkedListQueue) Enqueue(val int) {
	node := &Node{Val: val}
	if q.IsEmpty() {
		q.front = node
		q.rear = node
	} else {
		q.rear.Next = node
		q.rear = node
	}
	q.size++
}

func (q *LinkedListQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	val := q.front.Val
	q.front = q.front.Next
	q.size--
	if q.IsEmpty() {
		q.rear = nil
	}
	return val, true
}

func (q *LinkedListQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.front.Val, true
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Print() {
	fmt.Print("Queue (front -> rear): [")
	cur := q.front
	for cur != nil {
		fmt.Print(cur.Val)
		if cur.Next != nil {
			fmt.Print(", ")
		}
		cur = cur.Next
	}
	fmt.Println("]")
}

func Run() {
	q := &LinkedListQueue{}

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

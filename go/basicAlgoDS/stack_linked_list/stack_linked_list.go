package stack_linked_list

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedListStack struct {
	top  *Node
	size int
}

func (s *LinkedListStack) Push(val int) {
	node := &Node{Val: val, Next: s.top}
	s.top = node
	s.size++
}

func (s *LinkedListStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	val := s.top.Val
	s.top = s.top.Next
	s.size--
	return val, true
}

func (s *LinkedListStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.top.Val, true
}

func (s *LinkedListStack) Size() int {
	return s.size
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedListStack) Print() {
	fmt.Print("Stack (top -> bottom): [")
	cur := s.top
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
	s := &LinkedListStack{}

	fmt.Println("=== Push 1, 2, 3, 4, 5 ===")
	for _, v := range []int{1, 2, 3, 4, 5} {
		s.Push(v)
	}
	s.Print()
	fmt.Printf("Size: %d\n", s.Size())
	fmt.Printf("IsEmpty: %v\n", s.IsEmpty())

	if val, ok := s.Peek(); ok {
		fmt.Printf("\n=== Peek: %d ===\n", val)
	}
	s.Print()

	fmt.Println("\n=== Pop ===")
	if val, ok := s.Pop(); ok {
		fmt.Printf("Popped: %d\n", val)
	}
	s.Print()
	fmt.Printf("Size: %d\n", s.Size())

	fmt.Println("\n=== Pop all remaining ===")
	for !s.IsEmpty() {
		if val, ok := s.Pop(); ok {
			fmt.Printf("Popped: %d\n", val)
		}
	}
	s.Print()
	fmt.Printf("IsEmpty: %v\n", s.IsEmpty())

	fmt.Println("\n=== Pop from empty stack ===")
	_, ok := s.Pop()
	fmt.Printf("Pop success: %v\n", ok)
}

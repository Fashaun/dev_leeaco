package stack_array

import "fmt"

type ArrayStack struct {
	data []int
	size int
}

func (s *ArrayStack) Push(val int) {
	s.data = append(s.data, val)
	s.size++
}

func (s *ArrayStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	val := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return val, true
}

func (s *ArrayStack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.data[s.size-1], true
}

func (s *ArrayStack) Size() int {
	return s.size
}

func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

func (s *ArrayStack) Print() {
	fmt.Print("Stack (top -> bottom): [")
	for i := s.size - 1; i >= 0; i-- {
		fmt.Print(s.data[i])
		if i > 0 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func Run() {
	s := &ArrayStack{}

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

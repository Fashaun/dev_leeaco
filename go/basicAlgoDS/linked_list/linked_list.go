package linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

type LinkedList struct {
	Head *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) InsertAtIndex(index int, val int) {
	newNode := NewListNode(val)
	if index == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}
	curr := ll.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	newNode.Next = curr.Next
	curr.Next = newNode
}

func (ll *LinkedList) RemoveAtIndex(index int) {
	if index == 0 {
		ll.Head = ll.Head.Next
		return
	}
	curr := ll.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
}

func (ll *LinkedList) Access(index int) int {
	curr := ll.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr.Val
}

func (ll *LinkedList) Find(val int) int {
	curr := ll.Head
	index := 0
	for curr != nil {
		if curr.Val == val {
			return index
		}
		curr = curr.Next
		index++
	}
	return -1
}

func (ll *LinkedList) Print() {
	curr := ll.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func Run() {
	ll := NewLinkedList()

	// Initialize: insert values 1-5
	fmt.Println("=== Initialize Linked List ===")
	for i := 0; i < 5; i++ {
		ll.InsertAtIndex(i, i+1)
	}
	ll.Print()

	// Insert at head
	fmt.Println("\n=== Insert 0 at Head (index 0) ===")
	ll.InsertAtIndex(0, 0)
	ll.Print()

	// Insert in middle
	fmt.Println("\n=== Insert 10 at Index 3 ===")
	ll.InsertAtIndex(3, 10)
	ll.Print()

	// Access
	fmt.Println("\n=== Access Elements ===")
	fmt.Printf("Element at index 0: %d\n", ll.Access(0))
	fmt.Printf("Element at index 3: %d\n", ll.Access(3))

	// Find
	fmt.Println("\n=== Find Elements ===")
	fmt.Printf("Find 10: index %d\n", ll.Find(10))
	fmt.Printf("Find 99: index %d\n", ll.Find(99))

	// Remove head
	fmt.Println("\n=== Remove at Index 0 (head) ===")
	ll.RemoveAtIndex(0)
	ll.Print()

	// Remove middle
	fmt.Println("\n=== Remove at Index 2 (middle) ===")
	ll.RemoveAtIndex(2)
	ll.Print()
}

package deque_array

import "fmt"

type ArrayDeque struct {
	data []int
	size int
}

func (d *ArrayDeque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
	d.size++
}

func (d *ArrayDeque) PushRear(val int) {
	d.data = append(d.data, val)
	d.size++
}

func (d *ArrayDeque) PopFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	val := d.data[0]
	d.data = d.data[1:]
	d.size--
	return val, true
}

func (d *ArrayDeque) PopRear() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	val := d.data[d.size-1]
	d.data = d.data[:d.size-1]
	d.size--
	return val, true
}

func (d *ArrayDeque) PeekFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.data[0], true
}

func (d *ArrayDeque) PeekRear() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	return d.data[d.size-1], true
}

func (d *ArrayDeque) Size() int {
	return d.size
}

func (d *ArrayDeque) IsEmpty() bool {
	return d.size == 0
}

func (d *ArrayDeque) Print() {
	fmt.Print("Deque (front -> rear): [")
	for i := 0; i < d.size; i++ {
		fmt.Print(d.data[i])
		if i < d.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func Run() {
	d := &ArrayDeque{}

	fmt.Println("=== PushRear 1, 2, 3 ===")
	for _, v := range []int{1, 2, 3} {
		d.PushRear(v)
	}
	d.Print()

	fmt.Println("\n=== PushFront 0, -1 ===")
	d.PushFront(0)
	d.PushFront(-1)
	d.Print()
	fmt.Printf("Size: %d\n", d.Size())
	fmt.Printf("IsEmpty: %v\n", d.IsEmpty())

	if val, ok := d.PeekFront(); ok {
		fmt.Printf("\n=== PeekFront: %d ===\n", val)
	}
	if val, ok := d.PeekRear(); ok {
		fmt.Printf("=== PeekRear: %d ===\n", val)
	}
	d.Print()

	fmt.Println("\n=== PopFront ===")
	if val, ok := d.PopFront(); ok {
		fmt.Printf("PopFront: %d\n", val)
	}
	d.Print()

	fmt.Println("\n=== PopRear ===")
	if val, ok := d.PopRear(); ok {
		fmt.Printf("PopRear: %d\n", val)
	}
	d.Print()
	fmt.Printf("Size: %d\n", d.Size())

	fmt.Println("\n=== Pop all remaining ===")
	for !d.IsEmpty() {
		if val, ok := d.PopFront(); ok {
			fmt.Printf("PopFront: %d\n", val)
		}
	}
	d.Print()
	fmt.Printf("IsEmpty: %v\n", d.IsEmpty())

	fmt.Println("\n=== PopFront from empty deque ===")
	_, ok := d.PopFront()
	fmt.Printf("PopFront success: %v\n", ok)

	fmt.Println("\n=== PopRear from empty deque ===")
	_, ok = d.PopRear()
	fmt.Printf("PopRear success: %v\n", ok)
}

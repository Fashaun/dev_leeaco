package dynamic_array_list

import "fmt"

type DynamicArrayList struct {
	data     []int
	size     int
	capacity int
}

func NewDynamicArrayList(capacity int) *DynamicArrayList {
	return &DynamicArrayList{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (dal *DynamicArrayList) expand() {
	newCapacity := dal.capacity * 2
	newData := make([]int, newCapacity)
	copy(newData, dal.data[:dal.size])
	dal.data = newData
	fmt.Printf("  [Expand] capacity: %d -> %d\n", dal.capacity, newCapacity)
	dal.capacity = newCapacity
}

func (dal *DynamicArrayList) Append(val int) {
	if dal.size == dal.capacity {
		dal.expand()
	}
	dal.data[dal.size] = val
	dal.size++
}

func (dal *DynamicArrayList) Get(index int) int {
	return dal.data[index]
}

func (dal *DynamicArrayList) Insert(index int, val int) {
	if dal.size == dal.capacity {
		dal.expand()
	}
	for i := dal.size; i > index; i-- {
		dal.data[i] = dal.data[i-1]
	}
	dal.data[index] = val
	dal.size++
}

func (dal *DynamicArrayList) Delete(index int) {
	for i := index; i < dal.size-1; i++ {
		dal.data[i] = dal.data[i+1]
	}
	dal.size--
}

func (dal *DynamicArrayList) Traverse() {
	fmt.Printf("List (size=%d, capacity=%d): %v\n", dal.size, dal.capacity, dal.data[:dal.size])
}

func (dal *DynamicArrayList) Concat(other *DynamicArrayList) *DynamicArrayList {
	newCapacity := dal.size + other.size
	result := NewDynamicArrayList(newCapacity)
	copy(result.data, dal.data[:dal.size])
	copy(result.data[dal.size:], other.data[:other.size])
	result.size = dal.size + other.size
	return result
}

func (dal *DynamicArrayList) Sort() {
	for i := 1; i < dal.size; i++ {
		key := dal.data[i]
		j := i - 1
		for j >= 0 && dal.data[j] > key {
			dal.data[j+1] = dal.data[j]
			j--
		}
		dal.data[j+1] = key
	}
}

func Run() {
	// Initialize
	dal := NewDynamicArrayList(4)
	fmt.Println("=== Initialize Dynamic Array List (capacity=4) ===")
	dal.Traverse()

	// Append with expansion
	fmt.Println("\n=== Append Elements (showing expansion) ===")
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Printf("Appending %d...\n", v)
		dal.Append(v)
	}
	dal.Traverse()

	// Access
	fmt.Println("\n=== Access Elements ===")
	fmt.Printf("Element at index 0: %d\n", dal.Get(0))
	fmt.Printf("Element at index 3: %d\n", dal.Get(3))

	// Insert
	fmt.Println("\n=== Insert 10 at Index 2 ===")
	dal.Insert(2, 10)
	dal.Traverse()

	// Delete
	fmt.Println("\n=== Delete at Index 2 ===")
	dal.Delete(2)
	dal.Traverse()

	// Concat
	fmt.Println("\n=== Concatenate Two Lists ===")
	other := NewDynamicArrayList(4)
	for _, v := range []int{7, 8, 9} {
		other.Append(v)
	}
	fmt.Print("List A: ")
	dal.Traverse()
	fmt.Print("List B: ")
	other.Traverse()
	combined := dal.Concat(other)
	fmt.Print("Concat: ")
	combined.Traverse()

	// Sort
	fmt.Println("\n=== Sort (unsorted list) ===")
	unsorted := NewDynamicArrayList(8)
	for _, v := range []int{5, 2, 8, 1, 9, 3} {
		unsorted.Append(v)
	}
	fmt.Print("Before: ")
	unsorted.Traverse()
	unsorted.Sort()
	fmt.Print("After:  ")
	unsorted.Traverse()
}

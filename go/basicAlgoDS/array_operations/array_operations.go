package array_operations

import "fmt"

func InitArray(vals ...int) []int {
	return append([]int{}, vals...)
}

func AccessElement(arr []int, index int) int {
	return arr[index]
}

func InsertElement(arr []int, index int, val int) []int {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = val
	return arr
}

func RemoveElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func TraverseArray(arr []int) {
	for i, v := range arr {
		fmt.Printf("  Index %d: %d\n", i, v)
	}
}

func FindElement(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func ExpandArray(arr []int, vals ...int) []int {
	return append(arr, vals...)
}

func Run() {
	// 1. Initialize
	arr := InitArray(1, 2, 3, 4, 5)
	fmt.Println("=== Initialize Array ===")
	fmt.Println("Array:", arr)

	// 2. Access
	fmt.Println("\n=== Access Element ===")
	fmt.Printf("Element at index 2: %d\n", AccessElement(arr, 2))

	// 3. Insert
	fmt.Println("\n=== Insert Element ===")
	arr = InsertElement(arr, 2, 10)
	fmt.Println("After inserting 10 at index 2:", arr)

	// 4. Remove
	fmt.Println("\n=== Remove Element ===")
	arr = RemoveElement(arr, 2)
	fmt.Println("After removing element at index 2:", arr)

	// 5. Traverse
	fmt.Println("\n=== Traverse Array ===")
	TraverseArray(arr)

	// 6. Find
	fmt.Println("\n=== Find Element ===")
	fmt.Printf("Find 4: index %d\n", FindElement(arr, 4))
	fmt.Printf("Find 9: index %d\n", FindElement(arr, 9))

	// 7. Expand
	fmt.Println("\n=== Expand Array ===")
	arr = ExpandArray(arr, 6, 7, 8)
	fmt.Println("After expanding with [6, 7, 8]:", arr)
}

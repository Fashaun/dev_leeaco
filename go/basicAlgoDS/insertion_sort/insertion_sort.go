package insertion_sort

import "fmt"

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func Run() {
	arr := []int{5, 3, 8, 1, 2}
	fmt.Println("Before sorting:", arr)
	InsertionSort(arr)
	fmt.Println("After sorting:", arr)
}

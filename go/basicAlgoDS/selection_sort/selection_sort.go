package selection_sort

import "fmt"

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func Run() {
	arr := []int{5, 3, 8, 1, 2}
	fmt.Println("Before sorting:", arr)
	SelectionSort(arr)
	fmt.Println("After sorting:", arr)
}

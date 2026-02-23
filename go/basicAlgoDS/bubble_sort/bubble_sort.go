package bubble_sort

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Run() {
	arr := []int{5, 3, 8, 1, 2}
	fmt.Println("Before sorting:", arr)
	BubbleSort(arr)
	fmt.Println("After sorting:", arr)
}

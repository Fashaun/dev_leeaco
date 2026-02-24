package merge_sort

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		result := make([]int, len(arr))
		copy(result, arr)
		return result
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func Run() {
	arr := []int{5, 3, 8, 1, 2}
	fmt.Println("Before sorting:", arr)
	sorted := MergeSort(arr)
	fmt.Println("After sorting:", sorted)
}

## ADDED Requirements

### Requirement: partition rearranges elements around a pivot

The `quick_sort` package SHALL provide an unexported `partition(arr []int, low, high int) int` function that uses the Lomuto partition scheme with the last element (`arr[high]`) as the pivot. It SHALL rearrange elements so that all values less than or equal to the pivot come before it and all values greater than the pivot come after it, then return the final index of the pivot.

#### Scenario: Partition a subarray

- **WHEN** `partition([]int{3, 6, 8, 1, 5}, 0, 4)` is called (pivot=5)
- **THEN** elements ≤ 5 SHALL be moved before the pivot, elements > 5 after, and the function SHALL return the pivot's final index

#### Scenario: Partition a single element

- **WHEN** `partition(arr, i, i)` is called where `low == high`
- **THEN** the function SHALL return `low` without modifying the array

### Requirement: quickSort recursively sorts subarrays

The `quick_sort` package SHALL provide an unexported `quickSort(arr []int, low, high int)` function that recursively partitions and sorts subarrays. When `low < high`, it SHALL call `partition` to get the pivot index, then recursively sort the left subarray (`low` to `pivotIndex-1`) and the right subarray (`pivotIndex+1` to `high`).

#### Scenario: Recursive sort of unsorted array

- **WHEN** `quickSort([]int{5, 3, 8, 1, 2}, 0, 4)` is called
- **THEN** the array SHALL be sorted in-place to `[1, 2, 3, 5, 8]`

#### Scenario: Base case with single element

- **WHEN** `quickSort(arr, i, i)` is called where `low == high`
- **THEN** the function SHALL return without modifying the array

### Requirement: QuickSort sorts a slice in-place

The `quick_sort` package SHALL export a `QuickSort(arr []int)` function that sorts the given slice in ascending order in-place. It SHALL delegate to `quickSort(arr, 0, len(arr)-1)`.

#### Scenario: Sort an unsorted slice

- **WHEN** `QuickSort([]int{5, 3, 8, 1, 2})` is called
- **THEN** the slice SHALL be modified in-place to `[1, 2, 3, 5, 8]`

#### Scenario: Sort an already sorted slice

- **WHEN** `QuickSort([]int{1, 2, 3})` is called
- **THEN** the slice SHALL remain `[1, 2, 3]`

#### Scenario: Sort a single element slice

- **WHEN** `QuickSort([]int{42})` is called
- **THEN** the slice SHALL remain `[42]`

#### Scenario: Sort an empty slice

- **WHEN** `QuickSort([]int{})` is called
- **THEN** the function SHALL return without error

### Requirement: Run function displays quick sort demo

The `quick_sort` package SHALL export a `Run()` function that demonstrates quick sort by printing the array before and after sorting using the demo data `[5, 3, 8, 1, 2]`.

#### Scenario: Run is called

- **WHEN** `quick_sort.Run()` is invoked
- **THEN** the function SHALL print the array before sorting as `[5, 3, 8, 1, 2]` and after sorting as `[1, 2, 3, 5, 8]`

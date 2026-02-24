### Requirement: MergeSort function sorts a slice using divide-and-conquer

The `merge_sort` package SHALL provide a `MergeSort(arr []int) []int` function that returns a new sorted slice in ascending order using the recursive merge sort algorithm.

#### Scenario: Sort an unsorted slice

- **WHEN** `MergeSort([]int{5, 3, 8, 1, 2})` is called
- **THEN** the function SHALL return `[1, 2, 3, 5, 8]`

#### Scenario: Sort an already sorted slice

- **WHEN** `MergeSort` is called with a sorted slice `[1, 2, 3, 4, 5]`
- **THEN** the function SHALL return `[1, 2, 3, 4, 5]`

#### Scenario: Sort a single-element slice

- **WHEN** `MergeSort` is called with `[42]`
- **THEN** the function SHALL return `[42]`

#### Scenario: Sort an empty slice

- **WHEN** `MergeSort` is called with an empty slice
- **THEN** the function SHALL return an empty slice

#### Scenario: Original slice is not modified

- **WHEN** `MergeSort` is called with a slice
- **THEN** the original slice SHALL remain unchanged

### Requirement: merge helper combines two sorted slices

The `merge_sort` package SHALL provide an unexported `merge(left, right []int) []int` function that merges two sorted slices into a single sorted slice.

#### Scenario: Merge two sorted slices

- **WHEN** `merge([]int{1, 3, 5}, []int{2, 4, 6})` is called
- **THEN** the function SHALL return `[1, 2, 3, 4, 5, 6]`

#### Scenario: Merge with one empty slice

- **WHEN** `merge` is called with one empty slice and one non-empty sorted slice
- **THEN** the function SHALL return the non-empty slice's elements in order

### Requirement: Run function displays merge sort demo

The `merge_sort` package SHALL export a `Run()` function that demonstrates the merge sort algorithm with printed before/after output.

#### Scenario: Run is called

- **WHEN** `merge_sort.Run()` is invoked
- **THEN** the function SHALL print the array before sorting and the sorted result after calling `MergeSort`

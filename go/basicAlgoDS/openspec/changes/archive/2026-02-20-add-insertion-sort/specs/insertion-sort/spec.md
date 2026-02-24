## ADDED Requirements

### Requirement: InsertionSort function

The `insertion_sort` package SHALL provide an `InsertionSort(arr []int)` function that sorts the slice in-place in ascending order using the insertion sort algorithm.

#### Scenario: Sort an unsorted array

- **WHEN** `InsertionSort([]int{5, 3, 8, 1, 2})` is called
- **THEN** the slice SHALL be modified to `[1, 2, 3, 5, 8]`

#### Scenario: Sort an already sorted array

- **WHEN** `InsertionSort([]int{1, 2, 3, 4, 5})` is called
- **THEN** the slice SHALL remain `[1, 2, 3, 4, 5]`

#### Scenario: Sort a reverse sorted array

- **WHEN** `InsertionSort([]int{5, 4, 3, 2, 1})` is called
- **THEN** the slice SHALL be modified to `[1, 2, 3, 4, 5]`

#### Scenario: Sort a single element array

- **WHEN** `InsertionSort([]int{1})` is called
- **THEN** the slice SHALL remain `[1]`

#### Scenario: Sort an empty array

- **WHEN** `InsertionSort([]int{})` is called
- **THEN** the function SHALL complete without error

### Requirement: Run function displays sorting demo

The `insertion_sort` package SHALL export a `Run()` function that demonstrates the insertion sort algorithm with printed before and after output.

#### Scenario: Run is called

- **WHEN** `insertion_sort.Run()` is invoked
- **THEN** the function SHALL print the array before sorting and after sorting

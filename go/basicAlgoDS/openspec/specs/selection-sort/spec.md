### Requirement: SelectionSort function

The `selection_sort` package SHALL provide a `SelectionSort(arr []int)` function that sorts the slice in-place in ascending order using the selection sort algorithm.

#### Scenario: Sort an unsorted array

- **WHEN** `SelectionSort([]int{5, 3, 8, 1, 2})` is called
- **THEN** the slice SHALL be modified to `[1, 2, 3, 5, 8]`

#### Scenario: Sort an already sorted array

- **WHEN** `SelectionSort([]int{1, 2, 3, 4, 5})` is called
- **THEN** the slice SHALL remain `[1, 2, 3, 4, 5]`

#### Scenario: Sort a single element array

- **WHEN** `SelectionSort([]int{1})` is called
- **THEN** the slice SHALL remain `[1]`

#### Scenario: Sort an empty array

- **WHEN** `SelectionSort([]int{})` is called
- **THEN** the function SHALL complete without error

### Requirement: Run function displays sorting demo

The `selection_sort` package SHALL export a `Run()` function that demonstrates the selection sort algorithm with printed before and after output.

#### Scenario: Run is called

- **WHEN** `selection_sort.Run()` is invoked
- **THEN** the function SHALL print the array before sorting and after sorting

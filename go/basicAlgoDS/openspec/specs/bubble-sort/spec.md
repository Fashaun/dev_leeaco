### Requirement: Bubble sort algorithm implementation

The `bubble_sort` package SHALL export a `BubbleSort(arr []int)` function that sorts the given slice in-place in ascending order using the bubble sort algorithm.

#### Scenario: Sort an unsorted array

- **WHEN** `BubbleSort([]int{5, 3, 8, 1, 2})` is called
- **THEN** the slice SHALL be modified to `[1, 2, 3, 5, 8]`

#### Scenario: Sort an already sorted array

- **WHEN** `BubbleSort([]int{1, 2, 3})` is called
- **THEN** the slice SHALL remain `[1, 2, 3]`

#### Scenario: Sort an empty array

- **WHEN** `BubbleSort([]int{})` is called
- **THEN** the function SHALL complete without error

### Requirement: Run function displays sorting demo

The `bubble_sort` package SHALL export a `Run()` function that demonstrates the bubble sort algorithm by printing the array before and after sorting.

#### Scenario: Run is called

- **WHEN** `bubble_sort.Run()` is invoked
- **THEN** the function SHALL print the unsorted array and the sorted result to stdout

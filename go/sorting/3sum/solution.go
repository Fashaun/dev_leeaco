package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[i] + nums[lo] + nums[hi]
			if sum < 0 {
				lo++
			} else if sum > 0 {
				hi--
			} else {
				result = append(result, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				lo++
				hi--
			}
		}
	}

	return result
}

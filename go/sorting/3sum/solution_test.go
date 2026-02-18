package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func sortResult(result [][]int) {
	for _, triplet := range result {
		sort.Ints(triplet)
	}
	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]); k++ {
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return false
	})
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2 - no triplet",
			nums: []int{0, 1, 1},
			want: nil,
		},
		{
			name: "example 3 - all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			sortResult(got)
			sortResult(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

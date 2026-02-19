package cheapest_flights

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{
			name:    "path within stop limit",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0,
			dst:     3,
			k:       1,
			want:    700,
		},
		{
			name:    "cheaper one-stop route",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       1,
			want:    200,
		},
		{
			name:    "stop limit forces direct route",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       0,
			want:    500,
		},
		{
			name:    "no route exists",
			n:       3,
			flights: [][]int{{0, 1, 100}},
			src:     0,
			dst:     2,
			k:       1,
			want:    -1,
		},
		{
			name:    "no flights exist",
			n:       3,
			flights: [][]int{},
			src:     0,
			dst:     2,
			k:       1,
			want:    -1,
		},
		{
			name:    "single direct flight",
			n:       2,
			flights: [][]int{{0, 1, 100}},
			src:     0,
			dst:     1,
			k:       0,
			want:    100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			if got != tt.want {
				t.Errorf("findCheapestPrice(%d, %v, %d, %d, %d) = %d, want %d",
					tt.n, tt.flights, tt.src, tt.dst, tt.k, got, tt.want)
			}
		})
	}
}

package network_delay

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		times [][]int
		n, k  int
		want  int
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1},
	}
	for _, tt := range tests {
		got := networkDelayTime(tt.times, tt.n, tt.k)
		if got != tt.want {
			t.Errorf("networkDelayTime(%v, %d, %d) = %d, want %d", tt.times, tt.n, tt.k, got, tt.want)
		}
	}
}

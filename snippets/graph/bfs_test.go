package graph

import (
	"slices"
	"testing"
)

func TestBFS(t *testing.T) {
	// 0 - 1 - 2
	// |       |
	// 3       4     5 (孤立)
	adj := [][]int{
		0: {1, 3},
		1: {0, 2},
		2: {1, 4},
		3: {0},
		4: {2},
		5: {},
	}
	got := BFS(adj, 0)
	want := []int{0, 1, 2, 1, 3, -1}
	if !slices.Equal(got, want) {
		t.Errorf("BFS(adj, 0) = %v, want %v", got, want)
	}
}

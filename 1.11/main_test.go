package main

import (
	"testing"
)

func TestIntersectSets(t *testing.T) {
	tests := []struct {
		a, b     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{2, 3}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3}, []int{}, []int{}},
		{[]int{}, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 5, 9}, []int{2, 9, 10}, []int{9}},
		{[]int{1, 1, 2, 2, 3}, []int{2, 2, 3, 3}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
	}

	for _, tt := range tests {
		got := Intersect(tt.a, tt.b)
		if !equalIgnoreOrder(got, tt.expected) {
			t.Errorf("IntersectSets(%v, %v) = %v, want %v",
				tt.a, tt.b, got, tt.expected)
		}
	}
}

func equalIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[int]int)
	for _, v := range a {
		count[v]++
	}
	for _, v := range b {
		if count[v] == 0 {
			return false
		}
		count[v]--
	}
	return true
}

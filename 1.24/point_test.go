package main

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		p1, p2   *Point
		expected float64
	}{
		{NewPoint(0, 0), NewPoint(3, 4), 5},
		{NewPoint(1, 2), NewPoint(3, 4), math.Sqrt(8)},
		{NewPoint(-1, -1), NewPoint(1, 1), math.Sqrt(8)},
		{NewPoint(0, 0), NewPoint(0, 0), 0},
	}

	for _, tt := range tests {
		got := tt.p1.Distance(tt.p2)
		if math.Abs(got-tt.expected) > 1e-9 {
			t.Errorf("Distance(%v, %v) = %v, want %v", tt.p1, tt.p2, got, tt.expected)
		}
	}
}

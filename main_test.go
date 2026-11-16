package main

import (
	"math"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 1000} {
		data := generateRandomElements(size)
		if len(data) != max(size, 0) {
			t.Fatalf("size %d: got %d elements", size, len(data))
		}
		for i, value := range data {
			if value <= 0 {
				t.Fatalf("size %d: element %d is %d, want a positive integer", size, i, value)
			}
		}
	}
}

func TestMaximumSearch(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{math.MinInt}, math.MinInt},
		{"repeated maximum", []int{9, 2, 9, 1}, 9},
		{"eight chunks", []int{-3, -4, -2, -1, -7, -6, -9, -8}, -1},
		{"uneven chunks", []int{1, 8, 3, 4, 5, 6, 7, 2, 9, 10, math.MaxInt}, math.MaxInt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, want %d", got, tt.want)
			}
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, want %d", got, tt.want)
			}
		})
	}
}

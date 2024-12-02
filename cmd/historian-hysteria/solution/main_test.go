package main

import (
	"testing"
)

func TestTotalDistance(t *testing.T) {
	var tests = []struct {
		name  string
		left  []int
		right []int
		want  int
	}{
		{"site example", []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 11},
		{"identical lists", []int{1, 1, 1}, []int{1, 1, 1}, 0},
		{"empty lists", []int{}, []int{}, 0},
		{"different list length", []int{}, []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := TotalDistance(tt.left, tt.right)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestTotalSimilarityScore(t *testing.T) {
	var tests = []struct {
		name  string
		left  []int
		right []int
		want  int
	}{
		{"site example", []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 31},
		{"identical lists", []int{1, 1, 1}, []int{1, 1, 1}, 9},
		{"empty lists", []int{}, []int{}, 0},
		{"different list length", []int{}, []int{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := TotalSimilarityScore(tt.left, tt.right)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

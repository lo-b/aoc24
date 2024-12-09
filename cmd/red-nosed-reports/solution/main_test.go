package main

import (
	"testing"
)

// constants representing ordering
func TestSafeReports(t *testing.T) {
	var tests = []struct {
		name         string
		reportsArray [][]int
		want         int
	}{
		{
			"site example",
			[][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			2,
		},
		{
			"random rows from actual site input",
			[][]int{
				{27, 29, 32, 33, 36, 37, 40, 37},
				{58, 59, 58, 56, 53, 46, 41},
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

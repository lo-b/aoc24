package main

import (
	"testing"
)

func TestLevelValidatorCheck(t *testing.T) {
	const min = 1
	const max = 3
	var tests = []struct {
		name string
		pair [2]int
		sign int
		want bool
	}{
		{
			"AscendingInitiallyAscendingPair",
			[2]int{1, 2},
			-1,
			true,
		},
		{
			"AscendingInitiallyDescendingPair",
			[2]int{3, 1},
			-1,
			false,
		},
		{
			"AscendingInitiallyOutOfBounds",
			[2]int{3, 7},
			-1,
			false,
		},
		{
			"DescendingInitiallyOutOfBounds",
			[2]int{5, 0},
			1,
			false,
		},
		{
			"DescendingInitiallyEqualLevels",
			[2]int{5, 5},
			1,
			false,
		},
		{
			"AscendingInitiallyEqualLevels",
			[2]int{5, 5},
			-1,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := LevelValidator{sign: tt.sign, min: min, max: max}
			pairIsValid := validator.check(tt.pair[0], tt.pair[1])
			if pairIsValid != tt.want {
				t.Errorf("expected %v but got %v", tt.want, pairIsValid)
			}
		})
	}
}

func TestValidReportCheck(t *testing.T) {
	const min = 1
	const max = 3

	var tests = []struct {
		name   string
		levels []int
		want   bool
	}{
		{
			"SiteSample_#1_Valid",
			[]int{7, 6, 4, 2, 1},
			true,
		},
		{
			"SiteSample_#2_Invalid",
			[]int{1, 2, 7, 8, 9},
			false,
		},
		{
			"SiteSample_#3_Invalid",
			[]int{9, 7, 6, 2, 1},
			false,
		},
		{
			"SiteSample_#4_Invalid",
			[]int{1, 3, 2, 4, 5},
			false,
		},
		{
			"SiteSample_#5_Invalid",
			[]int{8, 6, 4, 4, 1},
			false,
		},
		{
			"SiteSample_#6_Valid",
			[]int{1, 3, 6, 7, 9},
			true,
		},
		{
			"SimpleAscendingLevels_IsValid",
			[]int{1, 2, 4},
			true,
		},
		{
			"SimpleDescendingLevels_IsValid",
			[]int{3, 2, -1},
			true,
		},
		{
			"HeadEqualsNext_IsInvalid",
			[]int{1, 1, 9},
			false,
		},
		{
			"TailEqualsPrev_IsInvalid",
			[]int{1, 9, 9},
			false,
		},
		{
			"SimpleAscendingLevels_InnerEqualLevels_IsInvalid",
			[]int{1, 2, 2, 3},
			false,
		},
		{
			"EqualLevels_IsInvalid",
			[]int{1, 1, 1, 1},
			false,
		},
		{
			"TwoAscendingElements_OutOfBounds_ReturnFalse",
			[]int{1, 9},
			false,
		},
		{
			"TwoDescendingElements_OutOfBounds_ReturnFalse",
			[]int{9, 1},
			false,
		},
		{
			"TwoDescendingElements_ReturnTrue",
			[]int{9, 6},
			true,
		},
		{
			"TwoAscendingElements_ReturnTrue",
			[]int{3, 6},
			true,
		},
		{
			"IdeanticalHeadTailLen3_ReturnFalse",
			[]int{3, 6, 3},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report := createReport(tt.levels, min, max)
			reportIsValid := report.isValid()
			if reportIsValid != tt.want {
				t.Errorf("expected %v but got %v", tt.want, reportIsValid)
			}
		})
	}
}

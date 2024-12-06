package main

import (
	"testing"

	dst "github.com/lo-b/aoc24/internal/datastructures"
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
			ans := SafeReports(constructReports(tt.reportsArray), false)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestReportIsSafe(t *testing.T) {
	var tests = []struct {
		name    string
		report  Report
		sorting int
		want    bool
	}{
		{
			name:    "single element report is valid due to recursion",
			report:  Report{dst.NewQueue(1)},
			sorting: None,
			want:    true,
		},
		{
			name:    "simple increasing list",
			report:  Report{dst.NewQueue(1, 2, 3, 4)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "simple decreasing list",
			report:  Report{dst.NewQueue(4, 3, 2, 1)},
			sorting: Desc,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ReportIsValid(tt.report.Queue.Head, tt.sorting, 1, 3)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestIsSafeWithValidation(t *testing.T) {
	var tests = []struct {
		name    string
		report  Report
		sorting int
		want    bool
	}{
		{
			// Safe without removing any level.
			name:    "Valid_SiteExampleReport_#1_DescendingLevels",
			report:  Report{dst.NewQueue(7, 6, 4, 2, 1)},
			sorting: Desc,
			want:    true,
		},
		{
			// Unsafe regardless of which level is removed.
			name:    "InValid_SiteExampleReport_#2_AscendingLevels_OutOfRange",
			report:  Report{dst.NewQueue(1, 2, 7, 8, 9)},
			sorting: Asc,
			want:    false,
		},
		{
			// Unsafe regardless of which level is removed.
			name:    "InValid_SiteExampleReport_#3_AscendingLevels_OutOfRange",
			report:  Report{dst.NewQueue(9, 7, 6, 2, 1)},
			sorting: Desc,
			want:    false,
		},
		{
			// Safe by removing the second level, 3.
			name:    "Valid_SiteExampleReport_#4_AscendingLevels_OutOfRange_ValidAfterRemovingSecondLevel",
			report:  Report{dst.NewQueue(1, 3, 2, 4, 5)},
			sorting: Asc,
			want:    true,
		},
		{
			// Safe by removing the third level, 4.
			name:    "Valid_SiteExampleReport_#5_AscendingLevels_Duplicates_ValidAfterRemovingThirdLevel",
			report:  Report{dst.NewQueue(8, 6, 4, 4, 1)},
			sorting: Desc,
			want:    true,
		},
		{
			// Safe without removing any level.
			name:    "Valid_SiteExampleReport_#6_AscendingLevels",
			report:  Report{dst.NewQueue(1, 3, 6, 7, 9)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "AscendingLevels_ValidAfterRemovingLevel",
			report:  Report{dst.NewQueue(1, 9, 2, 3, 4)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "Level_AscendingLevels_OutOfRange_ValidAfterRemovingLastLevel",
			report:  Report{dst.NewQueue(1, 2, 3, 4, 9)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "Valid_AscendingLevels_ValidAfterRemovingFirstLevel",
			report:  Report{dst.NewQueue(9, 1, 2, 3, 4)},
			sorting: Desc,
			want:    true,
		},
		{
			name:    "Valid_AscendingLevels_Duplicates_ValidAfterRemovingDuplicateLevel",
			report:  Report{dst.NewQueue(1, 1, 2, 3, 4, 5)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "Invalid_AscendingLevels_DuplicatesAndOutOfRange",
			report:  Report{dst.NewQueue(1, 1, 2, 6, 7, 8, 9)},
			sorting: Asc,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ReportIsValidWithToleration(tt.report.Queue.Head, tt.sorting, MinLevelDif, MaxLevelDiff, false)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func constructReports(reportsInput [][]int) []Report {
	var reports []Report
	for row := 0; row < len(reportsInput); row++ {
		nums := reportsInput[row]
		reports = append(reports, *CreateReport(nums...))
	}

	// Iterate over columns

	return reports
}

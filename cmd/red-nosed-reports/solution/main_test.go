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
			ans := SafeReportQueues(constructReports(tt.reportsArray), false)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestReportIsSafe(t *testing.T) {
	var tests = []struct {
		name    string
		report  ReportQueue
		sorting int
		want    bool
	}{
		{
			name:    "single element report is valid due to recursion",
			report:  ReportQueue{dst.NewQueue(1)},
			sorting: None,
			want:    true,
		},
		{
			name:    "simple increasing list",
			report:  ReportQueue{dst.NewQueue(1, 2, 3, 4)},
			sorting: Asc,
			want:    true,
		},
		{
			name:    "simple decreasing list",
			report:  ReportQueue{dst.NewQueue(4, 3, 2, 1)},
			sorting: Desc,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ValidQueue(tt.report.Queue.Head, tt.sorting, 1, 3)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestImprovedReportValid(t *testing.T) {
	var tests = []struct {
		name    string
		report  ReportDoublyLinkedList
		sorting int
		want    bool
	}{
		{
			// Safe without removing any level.
			name:    "Valid_SiteExampleReport_#1_DescendingLevels",
			report:  ReportDoublyLinkedList{dst.NewList(7, 6, 4, 2, 1)},
			sorting: 1,
			want:    true,
		},
		{
			// Unsafe regardless of which level is removed.
			name:    "InValid_SiteExampleReport_#2_AscendingLevels_OutOfRange",
			report:  ReportDoublyLinkedList{dst.NewList(1, 2, 7, 8, 9)},
			sorting: -1,
			want:    false,
		},
		{
			// Unsafe regardless of which level is removed.
			name:    "InValid_SiteExampleReport_#3_AscendingLevels_OutOfRange",
			report:  ReportDoublyLinkedList{dst.NewList(9, 7, 6, 2, 1)},
			sorting: 1,
			want:    false,
		},
		{
			// Safe by removing the second level, 3.
			name:    "Valid_SiteExampleReport_#4_AscendingLevels_OutOfRange_ValidAfterRemovingSecondLevel",
			report:  ReportDoublyLinkedList{dst.NewList(1, 3, 2, 4, 5)},
			sorting: -1,
			want:    true,
		},
		{
			// Safe by removing the third level, 4.
			name:    "Valid_SiteExampleReport_#5_AscendingLevels_Duplicates_ValidAfterRemovingThirdLevel",
			report:  ReportDoublyLinkedList{dst.NewList(8, 6, 4, 4, 1)},
			sorting: 1,
			want:    true,
		},
		{
			// Safe without removing any level.
			name:    "Valid_SiteExampleReport_#6_AscendingLevels",
			report:  ReportDoublyLinkedList{dst.NewList(1, 3, 6, 7, 9)},
			sorting: -1,
			want:    true,
		},
		{
			name:    "AscendingLevels_ValidAfterRemovingLevel",
			report:  ReportDoublyLinkedList{dst.NewList(1, 9, 2, 3, 4)},
			sorting: -1,
			want:    true,
		},
		{
			name:    "Level_AscendingLevels_OutOfRange_ValidAfterRemovingLastLevel",
			report:  ReportDoublyLinkedList{dst.NewList(1, 2, 3, 4, 9)},
			sorting: -1,
			want:    true,
		},
		{
			name:    "Valid_AscendingLevels_ValidAfterRemovingFirstLevel",
			report:  ReportDoublyLinkedList{dst.NewList(9, 1, 2, 3, 4)},
			sorting: 1,
			want:    true,
		},
		{
			name:    "Valid_AscendingLevels_Duplicates_ValidAfterRemovingDuplicateLevel",
			report:  ReportDoublyLinkedList{dst.NewList(1, 1, 2, 3, 4, 5)},
			sorting: 0,
			want:    true,
		},
		{
			name:    "Invalid_AscendingLevels_DuplicatesAndOutOfRange",
			report:  ReportDoublyLinkedList{dst.NewList(1, 1, 2, 6, 7, 8, 9)},
			sorting: 0,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ImprovedReportValid(tt.report.DoubleLinkedList, tt.report.DoubleLinkedList.Head, tt.sorting, MinLevelDif, MaxLevelDiff, false)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func constructReports(reportsInput [][]int) []ReportQueue {
	var reports []ReportQueue
	for row := 0; row < len(reportsInput); row++ {
		nums := reportsInput[row]
		reports = append(reports, *CreateReport(nums...))
	}

	return reports
}

func constructImprovedReports(reportsInput [][]int) []ReportDoublyLinkedList {
	var reports []ReportDoublyLinkedList
	for row := 0; row < len(reportsInput); row++ {
		nums := reportsInput[row]
		reports = append(reports, *CreateImprovedReport(nums...))
	}

	return reports
}

func TestIsValid_AscendingList(t *testing.T) {
	const min, max = 1, 3
	var tests = []struct {
		name     string
		list     *dst.DoubleLinkedList
		order    int
		expected bool
	}{
		{
			name:     "Valid ascending list (+)",
			list:     dst.NewList(1, 2, 3, 4),
			order:    -1,
			expected: true,
		},
		{
			name:     "Valid ascending list (-)",
			list:     dst.NewList(-4, -3, -2, -1),
			order:    -1,
			expected: true,
		},
		{
			name:     "Invalid list of identical (-) nums ",
			list:     dst.NewList(-4, -4, -4, -4),
			order:    0,
			expected: false,
		},
		{
			name:     "Invalid list of identical (+) nums ",
			list:     dst.NewList(4, 4, 4),
			order:    0,
			expected: false,
		},
		{
			name:     "Valid descending list (+)",
			list:     dst.NewList(4, 3, 2, -1),
			order:    1,
			expected: true,
		},
		{
			name:     "Valid descending list (-)",
			list:     dst.NewList(1, -2, -3, -4),
			order:    1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.list.Head
			for {
				if e.Next == nil {
					break
				}

				if IsValid(e, min, max, tt.order) != tt.expected {
					t.Errorf("expected element with key (%d) to be valid", e.Key)
				}
				e = e.Next
			}
		})
	}

}

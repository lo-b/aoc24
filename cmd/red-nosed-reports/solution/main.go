package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dst "github.com/lo-b/aoc24/internal/datastructures"
)

const (
	Asc          = iota // indicates ascending order.
	Desc         = iota // indicates descending order.
	None         = iota // indicates no order (unsorted).
	MinLevelDif  = 1    // minimum difference of adjacent levels.
	MaxLevelDiff = 3    // maximum adjacent level difference.
)

func main() {
	var useTolerance bool

	fmt.Println("Use tolerance module? true/False")
	fmt.Scanln(&useTolerance)

	file, err := os.Open("./assets/reports.txt")
	if err != nil {
		fmt.Println("Unable to read input file")
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v", err)
		}
	}()

	reader := bufio.NewReader(file)

	var reports []Report
	for {
		level, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		reportTxtVals := strings.Fields(level)
		var reportNums []int
		for _, reportTxtVal := range reportTxtVals {
			reportInt, _ := strconv.Atoi(reportTxtVal)

			reportNums = append(reportNums, reportInt)
		}

		reports = append(reports, *CreateReport(reportNums...))
	}

	fmt.Println("total valid report:", SafeReports(reports, useTolerance))
}

// Report encapsulates a report containing levels, which themselves are stored
// as numbers (satellite data) inside a Queue.
type Report struct {
	*dst.Queue
}

// SafeReports counts the total amount of reports which are safe.
func SafeReports(reports []Report, useTolerance bool) int {
	// total count of valid reports
	var validReportSum = 0
	for _, report := range reports {
		head, tail := report.Queue.Head, report.Queue.Tail
		var sorting int
		if head.Data < tail.Data {
			sorting = Asc
		} else if head.Data > tail.Data {
			sorting = Desc
		} else {
			sorting = None
		}

		var validReport bool
		if useTolerance {
			validReport = ReportIsValidWithToleration(head, sorting, MinLevelDif, MaxLevelDiff, false)

		} else {
			validReport = ReportIsValid(head, sorting, MinLevelDif, MaxLevelDiff)
		}

		if validReport {
			validReportSum++
		}
	}

	return validReportSum
}

// ReportIsValid checks validity of a report using the following criteria:
//   - levels of the report are either in ascending or descending order
//   - adjacent levels should differ least 1 and at most 3
//
// If both criteria are met return true, false otherwise.
func ReportIsValid(element *dst.Element, sorting int, min int, max int) bool {
	if element.Next == nil {
		return true
	}

	if sorting == Asc && element.Data >= element.Next.Data {
		return false
	}

	if sorting == Desc && element.Data <= element.Next.Data {
		return false
	}

	// Length that two adjacent levels differ by
	var adjacentLevelDiff int

	if sorting == Asc {
		adjacentLevelDiff = element.Next.Data - element.Data
	}

	if sorting == Desc {
		adjacentLevelDiff = element.Data - element.Next.Data
	}

	if adjacentLevelDiff < min || adjacentLevelDiff > max {
		return false
	}

	return ReportIsValid(element.Next, sorting, min, max)
}

func ReportIsValidWithToleration(element *dst.Element, sorting int, min int, max int, removedLevel bool) bool {
	if element.Next == nil {
		return true
	}

	// FIX: current implementation of tolerance will always skip the left
	// element (e.g. 9, 1, 2, 3) is valid but branching always keeps the
	// current element on branching and skipping next (recursive call with
	// .Next.Next). Bug above is prob fixed when also branching per discared
	// element, i.e.:
	//   - discard 1st => recursive call with .Next
	//   - discard 2nd => recursive call with .Next.Next (as is)

	if sorting == Asc && element.Data >= element.Next.Data {
		if !removedLevel && element.Next.Next == nil {
			return true
		}

		// calculate the new adjacent level diff of the next's next element
		// data, minus current element data
		var recalculatedDiff = element.Data - element.Next.Next.Data

		// reset the sorting
		if recalculatedDiff < 0 {
			sorting = Asc
		} else if recalculatedDiff > 0 {
			sorting = Desc
		} else {
			return false
		}

		var adjacentLevelDiff int
		if sorting == Asc {
			adjacentLevelDiff = element.Next.Next.Data - element.Data
		}

		if sorting == Desc {
			adjacentLevelDiff = element.Data - element.Next.Next.Data
		}

		inBounds := adjacentLevelDiff >= min && adjacentLevelDiff <= max

		if !removedLevel && inBounds && ReportIsValidWithToleration(element.Next.Next, sorting, min, max, true) {
			return true
		}

		return false
	}

	if sorting == Desc && element.Data <= element.Next.Data {
		if !removedLevel && element.Next.Next == nil {
			return true
		}

		// calculate the new adjacent level diff of the next's next element
		// data, minus current element data
		var recalculatedDiff = element.Data - element.Next.Next.Data

		// reset the sorting
		if recalculatedDiff < 0 {
			sorting = Asc
		} else if recalculatedDiff > 0 {
			sorting = Desc
		} else {
			return false
		}

		var adjacentLevelDiff int
		if sorting == Asc {
			adjacentLevelDiff = element.Next.Next.Data - element.Data
		}

		if sorting == Desc {
			adjacentLevelDiff = element.Data - element.Next.Next.Data
		}

		inBounds := adjacentLevelDiff >= min && adjacentLevelDiff <= max

		if !removedLevel && inBounds && ReportIsValidWithToleration(element.Next.Next, sorting, min, max, true) {
			return true
		}
		return false
	}

	// Length that two adjacent levels differ by
	var adjacentLevelDiff int

	if sorting == Asc {
		adjacentLevelDiff = element.Next.Data - element.Data
	}

	if sorting == Desc {
		adjacentLevelDiff = element.Data - element.Next.Data
	}

	if adjacentLevelDiff < min || adjacentLevelDiff > max {
		if element.Next.Next == nil {
			return true
		}

		// calculate the new adjacent level diff of the next's next element
		// data, minus current element data
		var recalculatedDiff = element.Data - element.Next.Next.Data

		// reset the sorting
		if recalculatedDiff < 0 {
			sorting = Asc
		}

		if recalculatedDiff > 0 {
			sorting = Desc
		}

		if recalculatedDiff == 0 {
			return false
		}

		if sorting == Asc {
			adjacentLevelDiff = element.Next.Next.Data - element.Data
		}

		if sorting == Desc {
			adjacentLevelDiff = element.Data - element.Next.Next.Data
		}

		inBounds := adjacentLevelDiff >= min && adjacentLevelDiff <= max

		if !removedLevel && inBounds && ReportIsValidWithToleration(element.Next.Next, sorting, min, max, true) {
			return true
		}
		return false
	}

	return ReportIsValidWithToleration(element.Next, sorting, min, max, false)
}

// CreateReport creates a new Report from nums integer arg(s).
func CreateReport(nums ...int) *Report {
	return &Report{
		Queue: dst.NewQueueFromArray(nums),
	}
}

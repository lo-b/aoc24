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
	minLevelDif  = 1    // minimum difference of adjacent levels.
	maxLevelDiff = 3    // maximum adjacent level difference.
)

func main() {
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

	fmt.Println("total valid report:", SafeReports(reports))
}

// Report encapsulates a report containing levels, which themselves are stored
// as numbers (satellite data) inside a Queue.
type Report struct {
	*dst.Queue
}

// SafeReports counts the total amount of reports which are safe.
func SafeReports(reports []Report) int {
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

		validReport := ReportIsValid(head, sorting, minLevelDif, maxLevelDiff)

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

// CreateReport creates a new Report from nums integer arg(s).
func CreateReport(nums ...int) *Report {
	return &Report{
		Queue: dst.NewQueueFromArray(nums),
	}
}

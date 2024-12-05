package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dst "github.com/lo-b/aoc24/internal/datastructures"
)

// constants representing ordering
const (
	// ascending order
	asc = iota
	// descending order
	desc = iota
	// no sorting
	none = iota

	min, max = 1, 3
)

func main() {
	file, err := os.Open("./assets/reports.txt")
	if err != nil {
		fmt.Println("Unable to read input file")
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	reports := []Report{}
	for {
		level, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		reportTxtVals := strings.Fields(level)
		reportNums := []int{}
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
			sorting = asc
		} else if head.Data > tail.Data {
			sorting = desc
		} else {
			sorting = none
		}

		validReport := ReportIsValid(head, sorting, min, max)

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

	if sorting == asc && element.Data >= element.Next.Data {
		return false
	}

	if sorting == desc && element.Data <= element.Next.Data {
		return false
	}

	// Length that two adjacent levels differ by
	var adjacentLevelDiff int

	if sorting == asc {
		adjacentLevelDiff = element.Next.Data - element.Data
	}

	if sorting == desc {
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

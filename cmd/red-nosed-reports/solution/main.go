package main

import (
	"bufio"
	"cmp"
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

	var reportLists []ReportDoublyLinkedList
	var reportQeueus []ReportQueue
	var validReportCount int
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

		if useTolerance {
			reportLists = append(reportLists, *CreateImprovedReport(reportNums...))
			// FIX: todo
		} else {
			reportQeueus = append(reportQeueus, *CreateReport(reportNums...))
		}
	}

	if !useTolerance {
		validReportCount = SafeReportQueues(reportQeueus, useTolerance)
	} else {
		// FIX: todo
		validReportCount = 0
	}

	fmt.Printf("total valid report: %d\n", validReportCount)
}

// ReportQueue encapsulates a report containing levels, which themselves are stored
// as numbers (satellite data) inside a Queue.
type ReportQueue struct {
	*dst.Queue
}

type ReportDoublyLinkedList struct {
	*dst.DoubleLinkedList
}

// SafeReportQueues counts the total amount of reports which are safe.
func SafeReportQueues(reports []ReportQueue, useTolerance bool) int {
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

		if ValidQueue(head, sorting, MinLevelDif, MaxLevelDiff) {
			validReportSum++
		}
	}

	return validReportSum
}

// ValidQueue checks validity of a report using the following criteria:
//   - levels of the report are either in ascending or descending order
//   - adjacent levels should differ least 1 and at most 3
//
// If both criteria are met return true, false otherwise.
func ValidQueue(element *dst.Element, sorting int, min int, max int) bool {
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

	return ValidQueue(element.Next, sorting, min, max)
}

func ImprovedReportValid(list *dst.DoubleLinkedList, element *dst.ListElement, sorting int, min int, max int, removedLevel bool) bool {
	// NOTE: not recovered previously and next element is Tail
	if !removedLevel && element.Next.Next == nil {
		return true
	}

	elementInvalid := !IsValid(element, min, max, sorting)

	// NOTE: element is invalid & previously recovered
	if elementInvalid && removedLevel {
		return false
	}

	if !elementInvalid && list.Head == list.Tail {
		return true
	}

	if !elementInvalid && list.Head == nil || list.Tail == nil {
		return true
	}
	if !elementInvalid && list.Head == list.Tail {
		return true
	}

	// try with recover
	if !removedLevel && elementInvalid {
		var listElementRemoved = CopyList(*list)
		var listChildRemoved = CopyList(*list)
		listElementRemoved.Delete(element)
		listChildRemoved.Delete(element.Next)
		return ImprovedReportValid(listElementRemoved, element.Next, sorting, min, max, true) ||
			ImprovedReportValid(listChildRemoved, element.Next.Next, sorting, min, max, true)
	}

	if !elementInvalid && element.Next != nil {
		list.Delete(element)
		return ImprovedReportValid(list, element.Next, sorting, min, max, removedLevel)
	}

	return false
}

func CopyList(list dst.DoubleLinkedList) *dst.DoubleLinkedList {
	e := list.Head
	listCopy := dst.NewEmptyList()
	for {
		if e.Next == nil {
			break
		}
		listCopy.Insert(&dst.ListElement{Key: e.Key})
		e = e.Next
	}

	return listCopy
}

func IsValid(element *dst.ListElement, min int, max int, sorting int) bool {
	var prevSorting int
	var prevLevelDiff int
	var nextSorting int
	var nextLevelDiff int

	if element.Next != nil {

		// -1 if x is less than y,
		//  0 if x equals y,
		// +1 if x is greater than y.
		nextSorting = cmp.Compare(element.Key, element.Next.Key)

		// NOTE: ascending order
		if nextSorting < 0 {
			nextLevelDiff = element.Next.Key - element.Key

			// NOTE: descending order
		} else if nextSorting > 0 {
			nextLevelDiff = element.Key - element.Next.Key
			// NOTE: equal curr & next element
		} else {
			return false
		}
	}

	if element.Prev != nil {
		// NOTE: compare elements 'left-to-right'
		prevSorting = cmp.Compare(element.Prev.Key, element.Key)

		// NOTE:
		if prevSorting < 0 {
			prevLevelDiff = element.Key - element.Prev.Key
			// NOTE: descending order
		} else if prevSorting > 0 {
			prevLevelDiff = element.Prev.Key - element.Key
			// NOTE: equal curr & next element
		} else {
			return false
		}

	}

	if element.Next != nil && (nextLevelDiff < min || nextLevelDiff > max || nextSorting != sorting) {
		return false
	}

	if element.Prev != nil && (prevLevelDiff < min || prevLevelDiff > max || prevSorting != sorting) {
		return false
	}

	return true
}

// CreateReport creates a new Report from nums integer arg(s).
func CreateReport(nums ...int) *ReportQueue {
	return &ReportQueue{
		Queue: dst.NewQueueFromArray(nums),
	}
}

// CreateReport creates a new Report from nums integer arg(s).
func CreateImprovedReport(nums ...int) *ReportDoublyLinkedList {
	return &ReportDoublyLinkedList{
		DoubleLinkedList: dst.NewListFromSlice(nums),
	}
}

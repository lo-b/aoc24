package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Asc          = -1 + iota // indicates ascending order.
	Desc         = iota      // indicates descending order.
	None         = iota      // indicates no order (unsorted).
	MinLevelDif  = 1         // minimum difference of adjacent levels.
	MaxLevelDiff = 3         // maximum adjacent level difference.
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

	var validReportCount = 0
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
			validReportCount = 1
		} else {
			validReportCount = -1
		}
	}

	fmt.Printf("total valid report: %d\n", validReportCount)
}

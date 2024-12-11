package main

import (
	"cmp"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/lo-b/aoc24/internal/puzzleio"
)

const (
	MinLevelDif  = 1 // minimum difference of adjacent levels.
	MaxLevelDiff = 3 // maximum adjacent level difference.
)

func main() {
	var useTolerance bool

	fmt.Println("Use tolerance module? true/False")
	fmt.Scanln(&useTolerance)

	puzzleInput, err := puzzleio.NewInputReader("./assets/reports.txt")
	file := puzzleInput.File
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

	reader := puzzleInput.Reader

	var validReportCount = 0
	for {
		level, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		reportTxtVals := strings.Fields(level)
		var levels []int
		for _, reportTxtVal := range reportTxtVals {
			reportInt, _ := strconv.Atoi(reportTxtVal)

			levels = append(levels, reportInt)
		}

		if useTolerance && validWithDampener(levels) {
			validReportCount++
		} else if !useTolerance {
			report := createReport(levels, MinLevelDif, MaxLevelDiff)
			if report.isValid() {
				validReportCount++
			}
		}
	}

	fmt.Printf("total valid report: %d\n", validReportCount)
}

// validWithDampener checks if levels are valid according to the following
// criteria:
//   - levels are either all increasing or all decreasing.
//   - two adjacent levels differ by at least one and at most three.
//   - tolerate a single bad level in what would otherwise be a safe report.
func validWithDampener(levels []int) bool {
	for k := 0; k < len(levels); k++ {
		var slicedLevels []int
		slicedLevels = append(slicedLevels, levels[:k]...)
		slicedLevels = append(slicedLevels, levels[k+1:]...)

		report := createReport(slicedLevels, MinLevelDif, MaxLevelDiff)
		if report.isValid() {
			return true
		}
	}

	return false
}

type Report struct {
	levels    []int
	validator Validator
}

type Validator interface {
	// check returns true if the integer pair x and y is valid, false
	// otherwise.
	check(x int, y int) bool
}

type LevelValidator struct {
	sign int
	min  int
	max  int
}

func createReport(levels []int, min int, max int) Report {
	sign := cmp.Compare(levels[0], levels[1])

	var report Report
	report.levels = levels
	report.validator = LevelValidator{sign, min, max}

	return report
}

// isValid checks if a Report is valid, according to the criteria:
//   - levels are either all increasing or all decreasing.
//   - two adjacent levels differ by at least one and at most three.
func (r Report) isValid() bool {
	validator, levels := r.validator, r.levels

	for levelIdx := range levels {
		hasNext := levelIdx <= len(levels)-2
		if hasNext {
			invalidPair := !validator.check(levels[levelIdx], levels[levelIdx+1])
			if invalidPair {
				return false
			}
		}
	}

	return true
}

// check determines whether levels x and y are valid, according to the
// following criteria:
//   - absolute difference of x and y is in the range [1,3]
//   - comparing x and y yields a sign that continues (previous)
//     ascending/descending order
func (validator LevelValidator) check(x int, y int) bool {

	newSign := cmp.Compare(x, y)
	var adjacentLevelDiff = int(math.Abs(float64(x - y)))
	if newSign != validator.sign ||
		adjacentLevelDiff < validator.min ||
		adjacentLevelDiff > validator.max {
		return false
	}

	return true
}

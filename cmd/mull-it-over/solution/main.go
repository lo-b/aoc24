package main

import (
	"fmt"
	"index/suffixarray"
	"strconv"
	"strings"

	"github.com/lo-b/aoc24/internal/puzzleio"
)

// NOTE: for mul to be 'valid' it has to be preceded by neither 'do' or 'don't'
// instructions or be preceded by the latest 'do' op which index is > don't
// op's index.
func main() {
	puzzleInput, err := puzzleio.NewPuzzleInput("assets/corrupted_memory_log.txt")
	if err != nil {
		fmt.Printf("Error reading puzzle input: %v", err)
		return
	}

	file := puzzleInput.File
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v", err)
			return
		}
	}()

	reader := puzzleInput.Reader

	var totalSum int
	var extendedTotalSum int
	var allLines string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		allLines += line
	}

	lineSum, extendedMulSum := Parse(allLines)
	totalSum = lineSum
	extendedTotalSum = extendedMulSum
	fmt.Printf("Total sum of 'mul' expressions: %d\n", totalSum)
	fmt.Printf("Total sum of 'mul' do/don't extended: %d\n", extendedTotalSum)
}

// Parse interprets valid multiplication expressions and returns their sum. A
// valid multiplication expression is defined as 'mul(X,Y)' where X, Y are ints
// in range [-999, 999].
func Parse(line string) (int, int) {
	const Expression = "mul("
	const Seperator = ','
	const EndChar = ')'

	indexedLine := suffixarray.New([]byte(line))
	// return all occurrences of valid expression start
	offsets := indexedLine.Lookup([]byte(Expression), -1)

	var mulSum int
	var extendedMulSum int
	for _, offset := range offsets {
		// mulIsEnabled := offsetEnabled(line, offset)
		mulSum += TryMulOperation(line, offset, Expression, Seperator, EndChar)

		if offsetEnabled(line, offset) {
			extendedMulSum += TryMulOperation(line, offset, Expression, Seperator, EndChar)
		}
	}

	return mulSum, extendedMulSum
}

func offsetEnabled(line string, offset int) bool {
	prevSlice := line[:offset]
	doIdx, dontIdx := strings.LastIndex(prevSlice, "do()"), strings.LastIndex(prevSlice, "don't()")

	// NOTE: no do/dont prev op found (standard enabled) => return true
	if doIdx == -1 && dontIdx == -1 {
		return true
	}

	// NOTE: no do(), don't() called before
	if doIdx == -1 && dontIdx > 0 {
		return false
	}

	if doIdx > dontIdx {
		return true
	}

	return false
}

func TryMulOperation(line string, offset int, operation string, seperator rune, endChar rune) int {
	digitRange := [2]int{-999, 999}
	lengthOffsetExpression := offset + len(operation)
	digitStrLength := max(len(strconv.Itoa(digitRange[0])), len(strconv.Itoa(digitRange[1])))
	sepIdx := strings.Index(line[offset:lengthOffsetExpression+digitStrLength], string(seperator))

	if sepIdx >= 0 {
		closeParenthesisIdx := strings.Index(line[offset:], string(endChar))
		possibleLeftDigitSlice := line[offset+len(operation) : offset+sepIdx]
		possibleRightDigitSlice := line[offset+sepIdx+1 : offset+closeParenthesisIdx]
		leftDigit, leftDigitErr := strconv.Atoi(possibleLeftDigitSlice)
		rightDigit, rightDigitErr := strconv.Atoi(possibleRightDigitSlice)

		if leftDigitErr != nil || rightDigitErr != nil {
			return 0
		}

		return leftDigit * rightDigit
	}

	return 0
}

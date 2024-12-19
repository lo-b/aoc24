package main

import (
	"fmt"
	"index/suffixarray"
	"io"
	"strconv"
	"strings"

	"github.com/lo-b/aoc24/internal/puzzleio"
)

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

	allLines, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	totalSum, extendedTotalSum := Parse(string(allLines))
	fmt.Printf("Total sum of 'mul' expressions: %d\n", totalSum)
	fmt.Printf("Total sum of 'mul' do/don't extended: %d\n", extendedTotalSum)
}

// Parse determines multiplication sums of valid 'mul' operations without and
// with conditional (do/don't) instructions, respectively.
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
		mulSum += TryMulOperation(line, offset, Expression, Seperator, EndChar)

		if OffsetEnabled(line, offset) {
			extendedMulSum += TryMulOperation(line, offset, Expression, Seperator, EndChar)
		}
	}

	return mulSum, extendedMulSum
}

// OffsetEnabled determines whether an instruction at a particular offset is
// enabled/disabled. Enabling or disabling instructions is done as follows:
//   - The do() instruction enables future instructions.
//   - The don't() instruction disables future instructions.
//
// Returns true when 'enabled', false otherwise.
func OffsetEnabled(line string, offset int) bool {
	prevSlice := line[:offset]
	doIdx, dontIdx := strings.LastIndex(prevSlice, "do()"), strings.LastIndex(prevSlice, "don't()")

	// NOTE: no previous do/don't instructions found; default should be 'enabled'
	if doIdx == -1 && dontIdx == -1 {
		return true
	}

	if doIdx == -1 && dontIdx > 0 {
		return false
	}

	if doIdx > dontIdx {
		return true
	}

	return false
}

// TryMulOperation tries to interpret a correctly starting mul operation --
// I.e. assumes the string slice starting at offset starts with the substring
// 'mul('. A valid multiplication expression is defined as 'mul(X,Y)' where X,
// Y are ints in range [-999, 999].
//
// Returns the multiplication result if operation is syntactically correct,
// return 0 in all other cases.
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

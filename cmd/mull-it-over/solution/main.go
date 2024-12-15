package main

import (
	"fmt"
	"index/suffixarray"
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

	reader := puzzleInput.Reader

	var totalSum int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lineSum := Parse(line)
		totalSum += lineSum
	}
	fmt.Printf("Total sum of 'mul' expressions: %d\n", totalSum)
}

// Parse searches line for valid multiplcation expressions, execute them, sum
// them up and finally return the total sum of a line. Instruction is defined
// as 'mul(X,Y)' where X, Y are ints in range [-999, 999].
func Parse(line string) int {
	const Instruction = "mul("
	const Seperator = ','
	const EndChar = ')'

	digitRange := [2]int{-999, 999}
	digitStrLength := max(len(strconv.Itoa(digitRange[0])), len(strconv.Itoa(digitRange[1])))
	indexedLine := suffixarray.New([]byte(line))

	// return all occurences of valid instruction start
	offsets := indexedLine.Lookup([]byte(Instruction), -1)

	var mulSum int
	for _, offset := range offsets {
		lengthOffsetInstruction := offset + len(Instruction)
		sepIdx := strings.Index(line[offset:lengthOffsetInstruction+digitStrLength], string(Seperator))

		if sepIdx >= 0 {
			closeParenthesisIdx := strings.Index(line[offset:], string(EndChar))
			possibleLeftDigitSlice := line[offset+len(Instruction) : offset+sepIdx]
			possibleRightDigitSlice := line[offset+sepIdx+1 : offset+closeParenthesisIdx]
			leftDigit, leftDigitErr := strconv.Atoi(possibleLeftDigitSlice)
			rightDigit, rightDigitErr := strconv.Atoi(possibleRightDigitSlice)

			if leftDigitErr != nil || rightDigitErr != nil {
				continue
			}

			mulSum += leftDigit * rightDigit
		} else {
			continue
		}
	}

	return mulSum
}

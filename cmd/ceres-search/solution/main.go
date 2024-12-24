package main

import (
	"fmt"

	"github.com/lo-b/aoc24/internal/puzzleio"
)

// Direction represents a compass direction as an enumerated integer type.
type Direction int

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func main() {
	puzzleInput, err := puzzleio.NewPuzzleInput("assets/puzzle.txt")
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

	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lines = append(lines, line)
	}

	inputSlices := CreateRuneGrid(lines)

	wordCount := WordSearch(inputSlices, "XMAS")
	xmasCount := XmasSearch(inputSlices)
	fmt.Printf("total word count: %d\n", wordCount)
	fmt.Printf("X-MAS occurences: %d\n", xmasCount)
}

// CreateRuneGrid creates a 2D-array of runes from a list of strings.
func CreateRuneGrid(stringInput []string) [][]rune {
	runeGrid := make([][]rune, 0, 0)
	for _, line := range stringInput {
		runeSlice := make([]rune, 0)
		for _, char := range line {
			runeSlice = append(runeSlice, char)
		}

		runeGrid = append(runeGrid, runeSlice)
	}

	return runeGrid
}

// XmasSearch finds the 'MAS' words in the shape of an X and returns the sum of total occurences.
func XmasSearch(puzzle [][]rune) int {
	var totalXmasCount int

	startingPoints := lookupLetter(puzzle, 'A')
	for _, point := range startingPoints {
		if validXmas(puzzle, point) {
			totalXmasCount++
		}
	}

	return totalXmasCount
}

// validXmas returns true if two 'MAS' words with shared 'A' (rune) are in the shape of an X.
func validXmas(puzzle [][]rune, point [2]int) bool {
	row, col := point[0], point[1]

	if row-1 < 0 || row+1 >= len(puzzle) {
		return false
	}

	if col-1 < 0 || col+1 >= len(puzzle[row]) {
		return false
	}

	if !xShaped(puzzle, point) {
		return false
	}

	return true
}

// xShaped returns true if 'MAS' words (forwards & backward) make an X shape.
func xShaped(puzzle [][]rune, point [2]int) bool {
	const word = "MAS"
	var wordCount int

	// loop over all possible diagonal directions (NE, SE, SW, NW)
	for d := 1; d <= 7; d += 2 {
		p := offSetPoint(Direction(d), point)
		diag := getDiag(Direction(d), p, puzzle, word)
		if string(diag) == word {
			wordCount++
		}

		if wordCount == 2 {
			return true
		}
	}

	return false
}

// offSetPoint returns a new point, offsetting point according to specified direction.
func offSetPoint(direction Direction, point [2]int) [2]int {
	row, col := point[0], point[1]
	if direction == NorthEast {
		return [2]int{row + 1, col - 1}
	} else if direction == SouthEast {
		return [2]int{row - 1, col - 1}
	} else if direction == SouthWest {
		return [2]int{row - 1, col + 1}
	} else if direction == NorthWest {
		return [2]int{row + 1, col + 1}
	}

	return point
}

// WordSearch looks in puzzle for occurrences of word and returns the sum of times word appears in the puzzle.
func WordSearch(puzzle [][]rune, word string) int {
	var totalWordCount int
	wordStartPoints := lookupLetter(puzzle, rune(word[0]))

	for _, point := range wordStartPoints {
		totalWordCount += walkPaths(point, word, puzzle)
	}

	return totalWordCount
}

// lookupLetter returns a list of points where letter occurs in puzzle.
func lookupLetter(puzzle [][]rune, letter rune) [][2]int {
	var points [][2]int
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == letter {
				points = append(points, [2]int{i, j})
			}
		}
	}

	return points
}

// walkPaths tries to walk all possible path directions and return the sum of paths containing the word to search for.
func walkPaths(point [2]int, word string, puzzle [][]rune) int {
	var (
		validWordCount int
		blockLen       = len(word) - 1
		row, col       = point[0], point[1]
	)

	if row-blockLen >= 0 && walk(North, point, puzzle, word) {
		validWordCount++
	}

	if row-blockLen >= 0 &&
		col+blockLen < len(puzzle[row]) &&
		walk(NorthEast, point, puzzle, word) {
		validWordCount++
	}

	if col+blockLen < len(puzzle[row]) && walk(East, point, puzzle, word) {
		validWordCount++
	}

	if row+blockLen < len(puzzle) &&
		col+blockLen < len(puzzle[row]) &&
		walk(SouthEast, point, puzzle, word) {
		validWordCount++
	}

	if row+blockLen < len(puzzle) && walk(South, point, puzzle, word) {
		validWordCount++
	}

	if col-blockLen >= 0 && row+blockLen < len(puzzle) &&
		walk(SouthWest, point, puzzle, word) {
		validWordCount++
	}

	if col-blockLen >= 0 && walk(West, point, puzzle, word) {
		validWordCount++
	}

	if col-blockLen >= 0 && row-blockLen >= 0 &&
		walk(NorthWest, point, puzzle, word) {
		validWordCount++
	}

	return validWordCount
}

// walk 'traverses' a path by getting the path as a rune slice and return true if
// the slice contains word, false otherwise.
func walk(direction Direction, point [2]int, puzzle [][]rune, word string) bool {
	if slice := pathSlice(direction, point, puzzle, word); string(slice) != word {
		return false
	}

	return true
}

// pathSlice constructs a slice of runes from a 'path', starting at point, going in the specified direction and having
// length equal to word.
func pathSlice(direction Direction, point [2]int, puzzle [][]rune, word string) []rune {
	if direction == East || direction == West {
		return getRow(direction, point, puzzle, word)
	}

	if direction == North || direction == South {
		return getCol(direction, point, puzzle, word)
	}

	return getDiag(direction, point, puzzle, word)
}

// getDiag returns a diagonal as a rune slice. It is created by starting from a point in the puzzle, going into
// direction and has a length equal to the word length.
func getDiag(direction Direction, point [2]int, puzzle [][]rune, word string) []rune {
	var diag []rune
	var x, y = point[0], point[1]
	if direction == NorthEast {
		for i := 0; i < len(word); i++ {
			diag = append(diag, puzzle[x-i][y+i])
		}
	} else if direction == SouthEast {
		for i := 0; i < len(word); i++ {
			diag = append(diag, puzzle[x+i][y+i])
		}
	} else if direction == SouthWest {
		for i := 0; i < len(word); i++ {
			diag = append(diag, puzzle[x+i][y-i])
		}
	} else if direction == NorthWest {
		for i := 0; i < len(word); i++ {
			diag = append(diag, puzzle[x-i][y-i])
		}
	} else {
		err := fmt.Errorf("invalid direction: %v", direction)
		println(err)
	}

	return diag
}

// getRow returns a row as a rune slice. The row is created by starting from a point in the puzzle, going into direction
// and has a length equal to the word length.
func getRow(direction Direction, point [2]int, puzzle [][]rune, word string) []rune {
	var row []rune
	var x, y = point[0], point[1]
	if direction == East {
		row = append(row, puzzle[x][y:y+len(word)]...)
	} else if direction == West {
		for i := 0; i < len(word); i++ {
			row = append(row, puzzle[x][y-i])
		}
	} else {
		err := fmt.Errorf("invalid direction: %v", direction)
		println(err)
	}

	return row
}

// getCol returns a column as a rune slice. The column is created by starting from a point in the puzzle, going into
// direction and has a length equal to the word length.
func getCol(direction Direction, point [2]int, puzzle [][]rune, word string) []rune {
	var col []rune
	var x, y = point[0], point[1]
	if direction == South {
		for i := 0; i < len(word); i++ {
			col = append(col, puzzle[x+i][y])
		}
	} else if direction == North {
		for i := 0; i < len(word); i++ {
			col = append(col, puzzle[x-i][y])
		}
	} else {
		err := fmt.Errorf("unexpected direction: %v", direction)
		println(err)
	}

	return col
}

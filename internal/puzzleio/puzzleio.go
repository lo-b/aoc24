package puzzleio

import (
	"bufio"
	"fmt"
	"os"
)

// PuzzleInput wraps bufio.Reader and os.File.
type PuzzleInput struct {
	Reader *bufio.Reader
	File   *os.File
}

// NewInputReader opens the file at the specified path and returns a
// PuzzleInput struct.
func NewInputReader(path string) (*PuzzleInput, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file: %w", err)
	}

	return &PuzzleInput{
		Reader: bufio.NewReader(file),
		File:   file,
	}, nil
}

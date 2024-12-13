package main

import (
	"fmt"

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

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		println(string(line))
	}
}

// Backus Naur Form (BNF) for 'mul(X,Y)' expression where X, Y are ints in
// range [-999, 999]:
// ----------------------------------------------------------------------------
// <program> ::= <instruction> | <instruction> <program>
// <instruction> ::= "mul" "(" <number> "," <number> ")"
// <number> ::= <digit> | <digit> <digit> | <digit> <digit> <digit>
// <digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
// ----------------------------------------------------------------------------
func parse() {
	panic("Not implemented")
}

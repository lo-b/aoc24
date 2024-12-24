package main

import "testing"

const word = "XMAS"

func TestWordSearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "Simple puzzle input containing word",
			input: []string{
				"XMAS",
			},
			expected: 1,
		},
		{
			name: "Simple puzzle input containing word with noise",
			input: []string{
				".XMASX",
			},
			expected: 1,
		},
		{
			name: "Puzzle containing word vertically & horizontally",
			input: []string{
				"XMAS.",
				"M....",
				"A....",
				"S....",
			},
			expected: 2,
		},
		{
			name: "Puzzle containing word two times horizontally & noise",
			input: []string{
				".XMAS",
				"XMAS.",
				"X.MAS",
				"XM.AS",
			},
			expected: 2,
		},
		{
			name: "Simple puzzle input containing word diagonally",
			input: []string{
				"X...",
				".M..",
				"..A.",
				"...S",
			},
			expected: 1,
		},
		{
			name: "Puzzle with word going in all directions (N, NE, E, SW, ...)",
			input: []string{
				"S..S..S",
				".A.A.A.",
				"..MMM..",
				"SAMXMAS",
				"..MMM..",
				".A.A.A.",
				"S..S..S",
			},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := WordSearch(CreateRuneGrid(tt.input), word)
			if actual != tt.expected {
				t.Errorf("Expected puzzle input\n%+v\nto contain %d occurrences of %s, found %d\n", tt.input, tt.expected, word, actual)
			}
		})
	}
}

func TestXmasSearch(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "Simple puzzle input containing single X-mas",
			input: []string{
				"S.M",
				".A.",
				"S.M",
			},
			expected: 1,
		},
		{
			name: "Site example",
			input: []string{
				".M.S......",
				"..A..MSMS.",
				".M.S.MAA..",
				"..A.ASMSM.",
				".M.S.M....",
				"..........",
				"S.S.S.S.S.",
				".A.A.A.A..",
				"M.M.M.M.M.",
				"..........",
			},
			expected: 9,
		},
		{
			name: "Puzzle input containing no X-mas",
			input: []string{
				"S..",
				".A.",
				"..M",
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := XmasSearch(CreateRuneGrid(tt.input))
			if actual != tt.expected {
				t.Errorf("Expected puzzle input\n%v to contain %d occurrences of X-mas, found %d\n.", tt.input, tt.expected, actual)
			}
		})
	}
}

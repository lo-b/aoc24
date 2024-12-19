package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		line        string
		expectedSum int
	}{
		{
			name:        "Simple line with single instruction",
			line:        "mul(2,4)",
			expectedSum: 8,
		},
		{
			name:        "Corrupted line with single instruction",
			line:        "@(*%$&)mul(5,7)...",
			expectedSum: 35,
		},
		{
			name:        "Corrupted line with multiple instructions",
			line:        "()@(*%$&)mul(5,7)...mul(5,2)",
			expectedSum: 45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSum, _ := Parse(tt.line)
			if actualSum != tt.expectedSum {
				t.Errorf("expected sum of mul instructions to be %v, got %v", tt.expectedSum, actualSum)
			}
		})
	}
}

func TestOffsetEnabled(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		offset   int
		expected bool
	}{
		{
			name:     "Simple line with single instruction",
			line:     "mul(2,4)",
			offset:   0,
			expected: true,
		},
		{
			name:     "Simple line with single instruction",
			line:     "don't()mul(2,4)",
			offset:   7,
			expected: false,
		},
		{
			name:     "Simple line with single instruction",
			line:     "do()...don't()mul(2,4)",
			offset:   15,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offsetIsEnabled := OffsetEnabled(tt.line, tt.offset)
			if offsetIsEnabled != tt.expected {
				t.Errorf("expected mul operation at offset :: %d :: to have 'enabled:%v', got %v", tt.offset, tt.expected, offsetIsEnabled)
			}
		})
	}
}

func TestExtendedMul(t *testing.T) {

	tests := []struct {
		name        string
		line        string
		expectedSum int
	}{
		{
			name:        "Part 2 site example",
			line:        "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,128](mul(11,8)undo()?mul(8,5))",
			expectedSum: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, extendedMulSum := Parse(tt.line)
			if extendedMulSum != tt.expectedSum {
				t.Errorf("expected sum of mul instructions to be %v, got %v", tt.expectedSum, extendedMulSum)
			}
		})
	}
}

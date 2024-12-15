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
			actualSum := Parse(tt.line)
			if actualSum != tt.expectedSum {
				t.Errorf("expected sum of mul instructions to be %v, got %v", tt.expectedSum, actualSum)
			}
		})
	}
}

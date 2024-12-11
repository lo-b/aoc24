package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/lo-b/aoc24/internal/puzzleio"
)

func main() {
	puzzleInput, err := puzzleio.NewInputReader("./assets/location_ids.txt")
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

	// left and right contain all location ids read of left and right col resp.
	var left, right []int
	lineNum := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		lineNum++
		values := strings.Fields(line)
		if len(values) != 2 {
			fmt.Printf("Input line %d does not contain a valid pair\n", lineNum)
			return
		}

		// WARN: Convert string to base-ten int. Ignore errors; appended 0 when
		// strconv fails.
		leftVal, _ := strconv.Atoi(values[0])
		rightVal, _ := strconv.Atoi(values[1])

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	fmt.Println("total distance:", TotalDistance(left, right))
	fmt.Println("total similarity score:", TotalSimilarityScore(left, right))
}

// TotalDistance calculates the sum of distancess between smallest pairs in
// left and right arrays.
func TotalDistance(left []int, right []int) int {
	leftSorted, rightSorted := make([]int, len(left)), make([]int, len(right))

	copy(leftSorted, left)
	copy(rightSorted, right)

	sort.Ints(leftSorted)
	sort.Ints(rightSorted)

	total := 0
	for idx, val := range leftSorted {
		total += int(math.Abs(float64(val - rightSorted[idx])))
	}

	return total
}

// TotalSimilarityScore calculates the total similarity score by summing values
// in the left list, each multiplied by the number of times it appears in the
// right list.
func TotalSimilarityScore(left []int, right []int) int {
	rightNumCounts := make(map[int]int)
	for _, num := range right {
		rightNumCounts[num] = rightNumCounts[num] + 1
	}

	total := 0
	for _, val := range left {
		total += val * rightNumCounts[val]
	}

	return total
}

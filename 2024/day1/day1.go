package main

import (
	"fmt"
	"sort"

	utility "example.com/go/aoc/utility"
)

func main() {
	lines := utility.ReadFileLineByLine("../../input/day1.txt")
	convertedLines := stringToIntArray(lines)

	sort.Ints(convertedLines[0])
	sort.Ints(convertedLines[1])

	d1a := distanceBetween(convertedLines)
	fmt.Printf("Day 1 Part 1: %d\n", d1a)
	d1b := similarityScore(convertedLines)
	fmt.Printf("Day 1 Part 2: %d\n", d1b)
}

func stringToIntArray(input []string) [][]int {
	output := make([][]int, 2)

	for _, val := range input {
		nums := utility.FetchSliceOfIntsInString(val)
		output[0] = append(output[0], nums[0])
		output[1] = append(output[1], nums[1])
	}

	return output
}

func distanceBetween(line [][]int) int {
	sum := 0
	length := len(line[0])

	for i := range length {
		diff := line[0][i] - line[1][i]
		if diff < 0 {
			diff *= -1
		}

		sum += diff
	}

	return sum
}

func similarityScore(line [][]int) int {
	score := 0
	length := len(line[0])

	m := make(map[int]int)

	for i := range length {
		m[line[1][i]] += 1
	}

	for _, key := range line[0] {
		if val, ok := m[key]; ok {
			score += key * val
		}
	}

	return score
}
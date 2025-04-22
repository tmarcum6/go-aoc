package main

import (
	"fmt"

	utility "example.com/go/aoc/utility"
)

func main() {
	lines := utility.ReadFileLineByLine("../../input/day2.txt")
	//convertedLines := stringToIntArray(lines)

	for l := range lines {
		fmt.Println(lines[l])
	} 

	//sort.Ints(convertedLines[0])
	//sort.Ints(convertedLines[1])

	//d1a := distanceBetween(convertedLines)
	//fmt.Printf("Day 1 Part 1: %d\n", d1a)
	//d1b := similarityScore(convertedLines)
	//fmt.Printf("Day 1 Part 2: %d\n", d1b)
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
package main

import (
	utility "example.com/go/aoc/utility"
	"fmt"
	"math"
)

func main() {
	lines := utility.ReadFileLineByLine("../../input/day2.txt")

	convertedLines := stringToIntSlice(lines)

	//d2a := distanceBetween(convertedLines)
	//fmt.Printf("Day 2 Part 1: %d\n", d2a)
	//d2b := similarityScore(convertedLines)
	//fmt.Printf("Day 2 Part 2: %d\n", d2b)
}

func stringToIntSlice(input []string) []int {
	var output []int
	var increase bool
	var decrease bool
	var safe bool
	var safeCount int
	var accetableDifference int = 3
	var unaccetable bool

	for _, val := range input {
		nums := utility.FetchSliceOfIntsInString(val)
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				decrease = true
				if nums[i-1]+nums[i] > accetableDifference {
					unaccetable = true
				}
			} else if nums[i-1] < nums[i] {
				increase = true
			}
		}

		if increase && decrease {
			safe = false
		}

		if safe {

			safeCount++
		}
	}

	return output
}

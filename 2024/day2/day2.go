package main

import (
	"fmt"

	utility "example.com/go/aoc/utility"
)

func main() {
	lines := utility.ReadFileLineByLine("../../input/day2.txt")
	safeCount, unsafeCount := safeCounter(lines)
	fmt.Printf("Day 2 Part 1 safe: %d\n", safeCount)
	fmt.Printf("Day 2 Part 1 unsafe: %d\n", unsafeCount)
}

func safeCounter(input []string) (int, int) {
	var safeCount int
	var unsafeCount int

	for _, val := range input {
		nums := utility.FetchSliceOfIntsInString(val)
		safe := checkSafety(nums)
		if safe {
			safeCount++
		} else {
			if problemDampener(nums) {
				safeCount++
			} else {
				unsafeCount++
			}
		}
	}

	return safeCount, unsafeCount
}

func checkSafety(input []int) bool {
	var increase bool = false
	var decrease bool = false
	var output bool = true

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]

		if diff > 0 {
			increase = true
		} else if diff < 0 {
			decrease = true
		} else {
			output = false
		}

		if diff > 3 || diff < -3 {
			output = false
		}

		if increase && decrease {
			output = false
		}
	}

	return output
}

func problemDampener(input []int) bool {
	var increase bool = false
	var decrease bool = false
	var output bool = false
	var unsafeCount int = 0

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]

		if diff > 0 {
			increase = true
		} else if diff < 0 {
			decrease = true
		} else {
			unsafeCount++
		}

		if diff > 3 || diff < -3 {
			unsafeCount++
		}

		if increase && decrease {
			unsafeCount++
		}
	}
	if unsafeCount == 1 {
		output = true
	}

	return output
}

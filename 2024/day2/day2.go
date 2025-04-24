package main

import (
	"fmt"

	utility "example.com/go/aoc/utility"
)

func main() {
	lines := utility.ReadFileLineByLine("../../input/day2.txt")
	safeCount, safeCountAfterRemove := safeCounter(lines)
	fmt.Printf("Safe Count: %d\n", safeCount)
	fmt.Printf("Safe Count After Remove: %d\n", safeCountAfterRemove)
	fmt.Printf("Total: %d\n", safeCount+safeCountAfterRemove)
}

func safeCounter(input []string) (int, int) {
	var safeCount int
	var safeCountAfterRemove int

	for _, val := range input {
		nums := utility.FetchSliceOfIntsInString(val)
		safe, i := checkSafety(nums)
		if safe {
			safeCount++
		} else if safeAfterRemoveBadIndex(nums, i) {
			safeCountAfterRemove++
		}
	}

	return safeCount, safeCountAfterRemove
}

func checkSafety(input []int) (bool, int) {
	var increase bool = false
	var decrease bool = false

	for i := 1; i < len(input); i++ {
		diff := input[i] - input[i-1]

		if diff > 0 {
			increase = true
		} else if diff < 0 {
			decrease = true
		} else {
			return false, i
		}

		if diff > 3 || diff < -3 {
			return false, i
		}

		if increase && decrease {
			return false, i
		}

	}

	return true, -1
}

func safeAfterRemoveBadIndex(input []int, badIndex int) bool {
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	inputCopy = utility.RemoveAt(inputCopy, badIndex)
	output, _ := checkSafety(inputCopy)

	if !output {
		fmt.Printf("Input Copy: %d\n", input)
		fmt.Printf("Input Copy: %d\n", inputCopy)
	}

	return output
}

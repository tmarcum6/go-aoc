package main

import (
	"fmt"
	"sort"

	utility "example.com/go/aoc/utility"
)

func main(){
	lines := utility.ReadFileLineByLine("C:/users/marcu/Documents/coding/go/aoc/input/day1.txt")
	convertedLines := stringToIntArray(lines)

	sort.Ints(convertedLines[0])
	sort.Ints(convertedLines[1])

	d1a := distanceBetween(convertedLines)
	fmt.Println(d1a)
	d1b := similarityScore(convertedLines)
	fmt.Println(d1b)
}

func stringToIntArray(s []string) [][]int{
	output := make([][]int, 2)

	for _, val := range s {
		nums := utility.FetchSliceOfIntsInString(val)
		output[0] = append(output[0], nums[0])
		output[1] = append(output[1], nums[1])
	}

	return output
}

func distanceBetween(line [][]int) int {

	return 1
}

func similarityScore(line [][]int) int {
	return 1
}

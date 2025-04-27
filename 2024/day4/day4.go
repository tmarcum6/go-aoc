package main

import (
	"fmt"
	"regexp"

	utility "example.com/go/aoc/utility"
)

// parse east
// parse west
// parse north
// parse south
// parse southeast
// parse southwest
// parse northwest
// parse northeast

func main() {
	input := utility.ReadFileLineByLine("../../input/day4.txt")
	re := regexp.MustCompile(`XMAS`)
	var matches [][]string
	count := 0
	for _, line := range input {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	fmt.Printf("count: %d\n", count)
}

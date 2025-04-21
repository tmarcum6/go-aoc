package main

import (
	utility "example.com/go/aoc/utility"
	"fmt"
)

func main() {
	lines := utility.ReadFileLineByLine("C:/users/marcu/Documents/coding/go/aoc/input/day1.txt")
	for _, line := range lines {
		if line == "" {
			continue
		}
		fmt.Println(line)
	}
	d1a := distanceBetween(lines)
	fmt.Println(d1a)
	d1b := similarityScore(lines)
	fmt.Println(d1b)
}

func distanceBetween(line []string) int {

	return 1
}

func similarityScore(line []string) int {
	return 1
}

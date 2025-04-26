package main

import (
	"fmt"
	"regexp"

	utility "example.com/go/aoc/utility"
)

func main() {
	input := utility.ReadFileLineByLine("../../input/day4.txt")
	re := regexp.MustCompile(`xmas`)
	var matches [][]string
	for _, line := range input {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}
	for _, m := range matches {
		fmt.Println(m)
	}

}

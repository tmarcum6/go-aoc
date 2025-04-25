package main

import (
	"fmt"
	"regexp"
	"strconv"

	utility "example.com/go/aoc/utility"
)

func main() {
	input := utility.ReadFileLineByLine("../../input/day3.txt")

	re := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`) // mul(x,y)

	var matches [][]string
	for _, line := range input {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	sum := 0
	for _, m := range matches {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		sum += x * y
		//fmt.Printf("mul(%d, %d) = %d\n", x, y, x*y)
	}
	fmt.Printf("sum: %d\n", sum)
}

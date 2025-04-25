package main

import (
	"fmt"
	"regexp"
	"strconv"

	utility "example.com/go/aoc/utility"
)

func main() {
	input := utility.ReadFileLineByLine("../../input/day3.txt")
	_ = findAllMatchesPart2(input)
	//sum := calculateSumOfMatches(matches)
	//fmt.Printf("sum: %d\n", sum)
}

func findAllMatchesPart1(input []string) [][]string {
	re := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`) // mul(x,y)

	var matches [][]string
	for _, line := range input {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
	}

	return matches
}

func findAllMatchesPart2(input []string) [][]string {
	re := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\)`)

	var matches [][]string
	for _, line := range input {
		matches = append(matches, re.FindAllStringSubmatch(line, -1)...)
		fmt.Printf("matches: %s\n", line)
	}

	return matches
}

func calculateSumOfMatches(matches [][]string) int {
	sum := 0
	for _, m := range matches {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		sum += x * y
		//fmt.Printf("mul(%d, %d) = %d\n", x, y, x*y)
	}
	return sum
}

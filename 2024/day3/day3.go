package main

import (
	"fmt"
	"regexp"
	"strconv"

	utility "example.com/go/aoc/utility"
)

func main() {
	input := utility.ReadFileLineByLine("../../input/day3.txt")

	matches1 := findAllMatchesPart1(input)
	sum1 := calculateSumOfMatches(matches1)
	fmt.Printf("sum part 1: %d\n", sum1)

	matches2 := findAllMatchesPart2(input)
	matches2 = filterMatches(matches2)
	sum2 := calculateSumOfMatches(matches2)
	fmt.Printf("sum part 2: %d\n", sum2)
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
	}

	return matches
}

func filterMatches(matches [][]string) [][]string {
	var filteredMatches [][]string

	p := 0
	for m := range matches {
		if matches[m][0] == "do()" {
			p = 1
		} else if matches[m][0] == "don't()" {
			p = 2
		}

		if p == 1 || p == 0 {
			filteredMatches = append(filteredMatches, matches[m])
		}
	}

	return filteredMatches
}

func calculateSumOfMatches(matches [][]string) int {
	sum := 0
	for _, m := range matches {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		sum += x * y
	}

	return sum
}

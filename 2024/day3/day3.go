package main

import (
	"fmt"
	"regexp"
	"strconv"

	utility "example.com/go/aoc/utility"
)

func main() {
	input := utility.ReadFileLineByLine("../../input/day3.txt")

	re := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`) // match mul(x,y)

	var matches [][]string
	matches = re.FindAllStringSubmatch(input[0], -1)

	if len(matches) == 0 {
		fmt.Println("no match")
		return
	}

	sum := 0
	for _, m := range matches {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])
		sum += x * y
		fmt.Printf("mul(%d, %d) = %d\n", x, y, x*y)
	}
	fmt.Printf("sum: %d\n", sum)
}

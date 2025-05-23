package aoc

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder

	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			nums = append(nums, int(localNum))
			build.Reset()
		}
	}

	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}

	return nums
}

func RemoveAt[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		return s
	}

	return slices.Delete(s, index, (index + 1))
}

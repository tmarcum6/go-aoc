package aoc

import(
		"strings"
		"strconv"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	for _, char := range line {
		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			nums = append(nums, int(localNum))
			build.Reset()
		}
	}
	return nums
}
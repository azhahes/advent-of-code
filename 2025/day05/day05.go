package day05

import (
	"bufio"
	"slices"
	"strings"

	"github.com/azhahes/advent-of-code/util/file"
	"github.com/azhahes/advent-of-code/util/parse"
)

func Part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	ranges := make([][2]int, 0)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		r := strings.Split(line, "-")
		ranges = append(ranges, [2]int{parse.MustToInt(r[0]), parse.MustToInt(r[1])})
	}

	for scanner.Scan() {
		id := parse.MustToInt(scanner.Text())
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				result++
				break
			}
		}
	}
	return result
}

func Part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	ranges := make([][2]int, 0)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		r := strings.Split(line, "-")
		ranges = append(ranges, [2]int{parse.MustToInt(r[0]), parse.MustToInt(r[1])})
	}
	slices.SortFunc(ranges, func(a, b [2]int) int { return a[0] - b[0] })
	newRanges := make([][2]int, 0)
	newRanges = append(newRanges, ranges[0])
	i := 0
	for _, r := range ranges {
		if newRanges[i][1] >= r[0] {
			newRanges[i][1] = max(r[1], newRanges[i][1])
		} else {
			newRanges = append(newRanges, r)
			i++
		}
	}
	for _, r := range newRanges {
		result += (r[1] - r[0] + 1)
	}
	return result
}

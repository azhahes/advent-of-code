package day2

import (
	"bufio"
	"strings"

	"github.com/azhahes/advent-of-code/2024/file"
	"github.com/azhahes/advent-of-code/2024/parse"
)

func part1(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	reports := make([][]int, 0)
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		reports = append(reports, parse.MustToIntSlice(levels))
	}
	var safeReports int
	for _, level := range reports {
		if len(level) == 1 {
			safeReports++
		}
		if checkDecreasing(level, 1, 3) || checkIncreasing(level, 1, 3) {
			safeReports++
		}
	}
	return safeReports
}

func checkIncreasing(level []int, min, max int) bool {
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if diff < min || diff > max {
			return false
		}
	}
	return true
}

func checkDecreasing(level []int, min, max int) bool {
	for i := 1; i < len(level); i++ {
		diff := level[i-1] - level[i]
		if diff < min || diff > max {
			return false
		}
	}
	return true
}

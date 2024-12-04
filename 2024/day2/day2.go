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
		if validateReport(level) {
			safeReports++
		}
	}
	return safeReports
}

func part2(filePath string) int {
	f := file.MustOpen(filePath)
	scanner := bufio.NewScanner(f)
	reports := make([][]int, 0)
	for scanner.Scan() {
		levels := strings.Fields(scanner.Text())
		reports = append(reports, parse.MustToIntSlice(levels))
	}
	var safeReports int
	for _, level := range reports {
		if validateReport(level) || checkWithFaultyLevel(level) {
			safeReports++
		}
	}
	return safeReports
}

func validateReport(report []int) bool {
	return checkOrder(report, 1, 3, func(a, b int) int { return b - a }) ||
		checkOrder(report, 1, 3, func(a, b int) int { return a - b })
}

func checkWithFaultyLevel(report []int) bool {
	for i := range len(report) {
		newReport := append([]int{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if validateReport(newReport) {
			return true
		}
	}
	return false
}

func checkOrder(level []int, min, max int, order func(int, int) int) bool {
	for i := 1; i < len(level); i++ {
		diff := order(level[i-1], level[i])
		if diff < min || diff > max {
			return false
		}
	}
	return true
}
